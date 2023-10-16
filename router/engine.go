package router

import "github.com/gin-gonic/gin"

func GetEngine() *gin.Engine {
	engine := gin.Default()
	//api
	ApiRouter(engine)
	//role
	RoleRouter(engine)
	//user
	UserRouter(engine)
	return engine
}
