package router

import (
	"github.com/gin-gonic/gin"
	"noopy-manager/controller"
)

func ApiRouter(engine *gin.Engine) {
	api := engine.Group("api")
	{
		api.POST("/create", controller.CreateApi)
		api.DELETE("/delete", controller.DeleteApi)
		api.PUT("/update", controller.UpdateApi)
		api.GET("/get", controller.GetApi)
	}
}
