package login

import (
	"github.com/gin-gonic/gin"
	v1 "xkginweb/api/v1"
)

// 登录路由
type LoginRouter struct{}

func (r *LoginRouter) InitLoginRouter(Router *gin.RouterGroup) {
	loginApi := v1.WebApiGroupApp.Login.LoginApi
	// 用组定义--（推荐）
	router := Router.Group("/login")
	{
		router.POST("/toLogin", loginApi.ToLogined)
	}
}
