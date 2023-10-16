package service

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"noopy-manager/global"
	"noopy-manager/model"
	"noopy-manager/utils"
)

func CreateApi(api model.Api) utils.Response {
	return utils.InsertOne(global.ApiColl, api)
}
func DeleteApi(_id primitive.ObjectID) utils.Response {
	//在角色中删除此权限
	if res, err := global.RoleColl.UpdateMany(context.TODO(), bson.M{}, bson.M{"$pull": bson.M{"apis": _id}}); err != nil {
		return utils.ErrorMess("更新角色", err.Error())
	} else {
		fmt.Println(res)
		return utils.DeleteOne(global.ApiColl, bson.M{"_id": _id})
	}
}
func UpdateApi(api model.Api) utils.Response {
	return utils.UpdateOne(global.ApiColl, bson.M{"_id": api.Id}, bson.M{"$set": api})
}
func GetApi(conditions map[string]interface{}, pageSize, currPage int64) utils.Response {
	return utils.GetPageData(global.ApiColl, conditions, pageSize, currPage)
}
