package dto

// admin注册请求
type RegisterRequest struct {
	Username        string `json:"username"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirmPassword"`
	Code            string `json:"code"`
	Uuid            string `json:"uuid"`
}

// admin登录请求
type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Code     string `json:"code"`
	Uuid     string `json:"uuid"`
}
