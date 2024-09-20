package login

import (
	"github.com/gin-gonic/gin"
	v1 "xkginweb/api/v1"
)

// 登录路由
type LogoutRouter struct{}

func (r *LogoutRouter) InitLogoutRouter(Router *gin.RouterGroup) {
	logoutApi := v1.WebApiGroupApp.Login.LogOutApi
	// 用组定义--（推荐）
	router := Router.Group("/login")
	{
		router.POST("/logout", logoutApi.ToLogout)
	}
}
