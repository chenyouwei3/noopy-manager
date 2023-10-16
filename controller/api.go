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

func CreateApi(c *gin.Context) {
	var api model.Api
	if err := c.Bind(&api); err != nil {
		c.JSON(http.StatusOK, utils.ErrorMess("参数错误", err.Error()))
		return
	}
	c.JSON(http.StatusOK, service.CreateApi(api))
}
func DeleteApi(c *gin.Context) {
	if _id, err := primitive.ObjectIDFromHex(c.Query("_id")); err != nil {
		c.JSON(http.StatusOK, utils.ErrorMess("参数错误", err.Error()))
		return
	} else {
		c.JSON(http.StatusOK, service.DeleteApi(_id))
	}
}
func UpdateApi(c *gin.Context) {
	var api model.Api
	if err := c.Bind(&api); err != nil {
		c.JSON(http.StatusOK, utils.ErrorMess("参数错误", err.Error()))
		return
	}
	c.JSON(http.StatusOK, service.UpdateApi(api))
}
func GetApi(c *gin.Context) {
	conditions := make(map[string]interface{})
	if method := c.Query("method"); method != "" {
		conditions["method"] = method
	}
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
	c.JSON(http.StatusOK, service.GetApi(conditions, pageSize, currPage))
}
