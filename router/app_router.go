package router

import (
	"template-project/app/controller"
	"template-project/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterAppGroupApi(api *gin.RouterGroup) {
	api.Use(middleware.Cors())

	api.POST("/register", (&controller.AuthController{}).AppRegister) // 注册
	api.POST("/login", (&controller.AuthController{}).AppLogin)       // 登录
	// 启用鉴权中间件，下面的路由都需要鉴权中间件验证通过后才可访问
	api.Use(middleware.AuthAppMiddleware())
}
