package routers

import (
	"gomemo/controller"

	"github.com/gin-gonic/gin"
)

func IndexRoutersInit(r *gin.Engine) {
	indexrouters := r.Group("/")
	{
		indexrouters.GET("/", controller.IndexController{}.Index)
		indexrouters.POST("/dopost", controller.IndexController{}.DoPost, redirecthtml)
		indexrouters.GET("/delete", controller.IndexController{}.DoDelete, redirecthtml)
		indexrouters.GET("/done", controller.IndexController{}.DoDone, redirecthtml)
		indexrouters.GET("/nodone", controller.IndexController{}.NoDone, redirecthtml)

	}
}

func redirecthtml(c *gin.Context) {
	c.Redirect(302, "/")
}
