package dto

// 用户授权
type UserTokenResponse struct {
	ID       int    `json:"ID"`
	UserName string `json:"userName"`
	NickName string `json:"nickName"`
	Password string `json:"-"`
	Status   string `json:"status"`
}
