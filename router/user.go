package router

import (
	"github.com/gin-gonic/gin"
	"noopy-manager/controller"
)

func UserRouter(engine *gin.Engine) {
	user := engine.Group("user")
	{
		user.POST("/create", controller.CreateUser)
		user.PUT("/update", controller.UpdateUser)
		user.GET("/get", controller.GetUser)
		user.DELETE("/delete", controller.DeleteUser)
	}
}
