package controller

import (
	"fmt"
	"strconv"
	"template-project/app/dto"
	"template-project/app/service"
	"template-project/app/token"
	"template-project/app/validator"
	"template-project/common/password"
	"template-project/common/types/constant"
	rediskey "template-project/common/types/redis-key"
	"template-project/config"
	"template-project/framework/dal"
	"template-project/framework/response"
	"time"

	"github.com/gin-gonic/gin"
)

type AuthController struct{}

// 退出登录
func (*AuthController) Logout(ctx *gin.Context) {

	token.DeleteToken(ctx)

	response.NewSuccess().Json(ctx)
}

// App注册
func (*AuthController) AppRegister(ctx *gin.Context) {

	var param dto.AppRegisterRequest

	if err := ctx.ShouldBind(&param); err != nil {
		response.NewError().SetMsg(err.Error()).Json(ctx)
		return
	}

	if err := validator.AppRegisterValidator(param); err != nil {
		response.NewError().SetMsg(err.Error()).Json(ctx)
		return
	}

	// if config := (&adminService.ConfigService{}).GetConfigCacheByConfigKey("sys.account.captchaEnabled"); config.ConfigValue == "true" {
	// 	if !captcha.NewCaptcha().Verify(param.Uuid, param.Code) {
	// 		response.NewError().SetMsg("验证码错误").Json(ctx)
	// 		return
	// 	}
	// }

	user := (&service.UserService{}).GetUserByUsername(param.Username)
	if user != nil && user.ID > 0 {
		response.NewError().SetMsg("保存用户" + param.Username + "注册账号已存在").Json(ctx)
		return
	}

	fmt.Println("注册用户名：", param.Username)

	//  创建用户
	if err := (&service.UserService{}).CreateUser(dto.User{
		Username: param.Username,
		Nickname: param.Username,
		Password: password.Generate(param.Password),
		Phone:    param.Username,
		Status:   1,
	}); err != nil {
		response.NewError().SetMsg(err.Error()).Json(ctx)
		return
	}

	response.NewSuccess().Json(ctx)
}

// app登录
func (*AuthController) AppLogin(ctx *gin.Context) {

	var param dto.LoginRequest

	if err := ctx.ShouldBind(&param); err != nil {
		response.NewError().SetCode(400).SetMsg(err.Error()).Json(ctx)
		return
	}

	if err := validator.AppLoginValidator(param); err != nil {
		response.NewError().SetCode(400).SetMsg(err.Error()).Json(ctx)
		return
	}

	user := (&service.UserService{}).GetUserByUsername(param.Username)
	if user.ID <= 0 || strconv.Itoa(user.Status) != constant.NORMAL_STATUS {
		response.NewError().SetMsg("用户不存在或被禁用").Json(ctx)
		return
	}

	// 登陆密码错误次数超过限制，锁定账号10分钟
	count, _ := dal.Redis.Get(ctx.Request.Context(), rediskey.LoginPasswordErrorKey+param.Username).Int()
	if count >= config.Data.User.Password.MaxRetryCount {
		response.NewError().SetMsg("密码错误次数超过限制，请" + strconv.Itoa(config.Data.User.Password.LockTime) + "分钟后重试").Json(ctx)
		return
	}

	if !password.Verify(user.Password, param.Password) {
		// 密码错误次数加1，并设置缓存过期时间为锁定时间
		dal.Redis.Set(ctx.Request.Context(), rediskey.LoginPasswordErrorKey+param.Username, count+1, time.Minute*time.Duration(config.Data.User.Password.LockTime))
		response.NewError().SetMsg("密码错误").Json(ctx)
		return
	}

	// 登录成功，删除错误次数
	dal.Redis.Del(ctx.Request.Context(), rediskey.LoginPasswordErrorKey+param.Username)

	userToken := dto.UserTokenResponse{
		ID:       user.ID,
		UserName: user.Username,
		NickName: user.Nickname,
		Password: user.Password,
		Status:   strconv.Itoa(user.Status),
	}

	tokenStr, err := token.GetAppClaims().GenerateAPPToken(userToken)
	if err != nil {
		response.NewError().SetMsg(err.Error()).Json(ctx)
		return
	}

	// // 更新登录的ip和时间
	(&service.UserService{}).UpdateUser(dto.User{
		ID:            user.ID,
		LastLoginIP:   ctx.ClientIP(),
		LastLoginTime: time.Now(),
	})

	response.NewSuccess().SetData("token", tokenStr).Json(ctx)
}
