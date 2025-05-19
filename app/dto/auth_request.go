package dto

// app注册请求
type AppRegisterRequest struct {
	Username        string `json:"username"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirmPassword"`
	Code            string `json:"code"`
}
