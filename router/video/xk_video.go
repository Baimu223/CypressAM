package video

import (
	"github.com/gin-gonic/gin"
	v1 "xkginweb/api/v1"
)

type XkVideoRouter struct{}

func (r *XkVideoRouter) InitXkVideoRouter(Router *gin.RouterGroup) {

	xkVideoApi := v1.WebApiGroupApp.Video.XkVideoApi
	// 这个路由多了一个对对post，put请求的中间件处理，而这个中间件做了一些对post和put的参数的处理和一些公共信息的处理
	router := Router.Group("video") //.Use(middleware.OperationRecord())
	{
		// 查询课程
		router.GET("find/page", xkVideoApi.FindVideosPage)
		router.GET("get/:id", xkVideoApi.GetVideosById)
		router.POST("save", xkVideoApi.SaveVideo)
		router.POST("update", xkVideoApi.UpdateVideo)

	}
}
