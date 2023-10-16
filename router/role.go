package router

import (
	"github.com/gin-gonic/gin"
	"noopy-manager/controller"
)

func RoleRouter(engine *gin.Engine) {
	role := engine.Group("role")
	{
		role.POST("/create", controller.CreateRole)
		role.PUT("/update", controller.UpdateRole)
		role.GET("/get", controller.GetRole)
	}
}
