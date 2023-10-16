package controller

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"noopy-manager/model"
	"noopy-manager/service"
	"noopy-manager/utils"
	"strconv"
)

func CreateRole(c *gin.Context) {
	var role model.Role
	if err := c.Bind(&role); err != nil {
		c.JSON(http.StatusOK, utils.ErrorMess("参数错误", err.Error()))
		return
	}
	c.JSON(http.StatusOK, service.CreateRole(role))
}
func UpdateRole(c *gin.Context) {
	var role model.Role
	if err := c.Bind(&role); err != nil {
		c.JSON(http.StatusOK, utils.ErrorMess("参数错误", err.Error()))
		return
	}
	c.JSON(http.StatusOK, service.UpdateRole(role))
}
func GetRole(c *gin.Context) {
	conditions := make(map[string]interface{})
	if name := c.Query("name"); name != "" {
		//i忽略大小写
		conditions["name"] = primitive.Regex{Pattern: name, Options: "i"}
	}
	//默认获取全部数据
	pageSize, err := strconv.ParseInt(c.DefaultQuery("pageSize", "0"), 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, utils.ErrorMess(err.Error(), nil))
		return
	}
	currPage, err := strconv.ParseInt(c.DefaultQuery("currPage", "1"), 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, utils.ErrorMess(err.Error(), nil))
		return
	}
	c.JSON(http.StatusOK, service.GetRole(conditions, pageSize, currPage))
}
