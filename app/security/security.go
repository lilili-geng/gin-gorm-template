package security

import (
	"template-project/app/token"

	"github.com/gin-gonic/gin"
)

// 获取用户id
//
// 例如：security.GetAuthUserId(ctx)
func GetAuthUserId(ctx *gin.Context) int {

	authUser, err := token.GetAuhtUser(ctx)
	if err != nil {
		return 0
	}

	return authUser.ID
}

// 获取用户账户
//
// 例如：security.GetAuthUserName(ctx)
func GetAuthUserName(ctx *gin.Context) string {

	authUser, err := token.GetAuhtUser(ctx)
	if err != nil {
		return ""
	}

	return authUser.UserName
}

// 获取用户
//
// 例如：security.GetAuthUser(ctx)
func GetAuthUser(ctx *gin.Context) *token.UserTokenResponse {

	authUser, err := token.GetAuhtUser(ctx)
	if err != nil {
		return nil
	}

	return authUser
}

// 获取app用户
//
// 例如：security.GetAuthUser(ctx)
func GetAppAuthUser(ctx *gin.Context) *token.UserTokenResponse {

	authUser, err := token.GetAppAuhtUser(ctx)
	if err != nil {
		return nil
	}

	return authUser
}
