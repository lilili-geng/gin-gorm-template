package validator

import (
	"errors"
	"template-project/app/dto"
)

// App注册验证
func AppRegisterValidator(param dto.AppRegisterRequest) error {

	if param.Username == "" {
		return errors.New("用户名不能为空")
	}

	if param.Password == "" {
		return errors.New("密码不能为空")
	}

	if param.ConfirmPassword != param.Password {
		return errors.New("两次密码不一致")
	}

	if param.Code == "" {
		return errors.New("验证码不能为空")
	}

	if len(param.Password) < 5 || len(param.Password) > 20 {
		return errors.New("密码长度必须在5-20之间")
	}

	return nil
}

// admin登录验证
func AppLoginValidator(param dto.LoginRequest) error {

	if param.Username == "" {
		return errors.New("用户名不能为空")
	}

	if param.Password == "" {
		return errors.New("密码不能为空")
	}

	return nil
}
