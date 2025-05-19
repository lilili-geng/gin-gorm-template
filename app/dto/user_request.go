package dto

import "time"

// 保存用户
type User struct {
	ID            int       `json:"id"`
	Username      string    `json:"username"`
	Password      string    `json:"password,omitempty"`
	Phone         string    `json:"phone"`
	Email         string    `json:"email"`
	Nickname      string    `json:"nickname"`
	Avatar        string    `json:"avatar"`
	RealName      string    `json:"realName"`
	Gender        int       `json:"gender"`
	Status        int       `json:"status"`
	LastLoginTime time.Time `json:"last_login_time"`
	LastLoginIP   string    `json:"last_login_ip"`
}

// admin登录请求
type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Code     string `json:"code"`
}
