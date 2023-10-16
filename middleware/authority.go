package middleware

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"net/http"
	"noopy-manager/global"
	"noopy-manager/model"
	"noopy-manager/utils"
)

func ApiAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		//获取访问api的url和方法
		method, url := c.Request.Method, c.Request.URL.Path
		fmt.Println(url)
		var api model.Api
		if err := global.ApiColl.FindOne(context.TODO(), bson.M{"url": url, "method": method}).Decode(&api); err != nil {
			c.JSON(http.StatusOK, utils.ErrorMess("验证api：此api不存在", err.Error()))
			c.Abort()
			return
		}
		//获取token解析出来的user
		userInterface, _ := c.Get("user")
		user := userInterface.(model.User)
		//获取user对应的role
		var role model.Role
		if err := global.RoleColl.FindOne(context.TODO(), bson.M{"_id": user.RoleId}).Decode(&role); err != nil {
			c.JSON(http.StatusOK, utils.ErrorMess("验证api：获取用户角色失败", err.Error()))
			c.Abort()
			return
		}
		//轮询role对应的apis，判断其是否相应的权限
		for i := range role.Apis {
			if role.Apis[i] == api.Id {
				c.Next()
				return
			}
		}
		c.JSON(http.StatusOK, utils.ErrorMess("验证api：此用户无访问此api的权限", nil))
		c.Abort()
		return
	}
}
