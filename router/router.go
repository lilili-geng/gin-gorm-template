package router

import (
	"github.com/gin-gonic/gin"
)

// 注册路由
func Register(server *gin.Engine) {

	api := server.Group("/api")
	app := server.Group("/app")

	RegisterAppGroupApi(app)

	RegisterAdminGroupApi(api)

}
