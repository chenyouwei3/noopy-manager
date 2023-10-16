package service

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/bcrypt"
	"math/rand"
	"noopy-manager/global"
	"noopy-manager/middleware"
	"noopy-manager/model"
	"noopy-manager/utils"
	"strconv"
	"time"
)

func Login(user model.User) utils.Response {
	var DBUser model.User
	//校验账号
	if err := global.UserColl.FindOne(context.TODO(), bson.M{"account": user.Account}).Decode(&DBUser); err != nil {
		return utils.ErrorMess("账号错误", err.Error())
	}
	//校验密码
	if err := bcrypt.CompareHashAndPassword([]byte(DBUser.Password), []byte(user.Password+DBUser.Salt)); err != nil {
		return utils.ErrorMess("密码错误", err.Error())
	}
	//查询角色信息
	var role model.Role
	if err := global.RoleColl.FindOne(context.TODO(), bson.M{"_id": DBUser.RoleId}).Decode(&role); err != nil {
		return utils.ErrorMess("此用户的角色不存在", err.Error())
	}
	//生成token
	token, err := middleware.CreateToken(DBUser)
	if err != nil {
		return utils.ErrorMess("生成token失败", err.Error())
	}
	res := map[string]interface{}{
		"_id":        DBUser.Id,
		"account":    DBUser.Account,
		"password":   DBUser.Password,
		"name":       DBUser.Name,
		"sex":        DBUser.Sex,
		"phone":      DBUser.Phone,
		"role":       role.Name,
		"roleCode":   role.Code,
		"AvatarUrl":  DBUser.AvatarUrl,
		"token":      token,
		"roleRoutes": role.RoleRoutes,
		"firstPage":  role.FirstPage,
	}
	return utils.SuccessMess("登陆成功", res)
}

func CreateUser(user model.User) utils.Response {
	//判断账号是否重复
	if err := global.UserColl.FindOne(context.TODO(), bson.M{"account": user.Account}).Decode(&bson.M{}); err == mongo.ErrNoDocuments {
		//判断角色是否存在
		if err = global.RoleColl.FindOne(context.TODO(), bson.M{"_id": user.RoleId}).Decode(&bson.M{}); err == mongo.ErrNoDocuments {
			return utils.ErrorMess("角色不存在", err)
		}
		rand.Seed(time.Now().Unix()) //根据时间戳生成种子
		//生成盐
		salt := strconv.FormatInt(rand.Int63(), 10)
		//密码加盐加密
		encryptedPass, err := bcrypt.GenerateFromPassword([]byte(user.Password+salt), bcrypt.DefaultCost)
		if err != nil {
			return utils.ErrorMess("密码加密失败", err)
		}
		user.Password, user.Salt = string(encryptedPass), salt
		//存入数据库
		return utils.InsertOne(global.UserColl, user)
	} else {
		return utils.ErrorMess("账号重复", err)
	}
}

func DeleteUser(_id primitive.ObjectID) utils.Response {
	return utils.DeleteOne(global.UserColl, bson.M{"_id": _id})
}

func UpdateUser(user model.User) utils.Response {
	//校验角色是否存在
	if err := global.RoleColl.FindOne(context.TODO(), bson.M{"_id": user.RoleId}).Decode(&bson.M{}); err != nil {
		return utils.ErrorMess(err.Error()+"此角色不存在", user)
	}
	return utils.UpdateOne(global.UserColl, bson.M{"_id": user.Id}, bson.M{"$set": user})
}

func GetUser(conditions map[string]interface{}, pageSize, currPage int64) utils.Response {
	var pageData []map[string]interface{}
	skip := (currPage - 1) * pageSize
	//获取分页数据
	if err := utils.Find(global.UserColl, &pageData, conditions, &options.FindOptions{
		Limit: &pageSize,
		Skip:  &skip,
		Sort:  bson.M{"_id": -1},
	}); err != nil {
		return utils.ErrorMess("获取分页数据失败", err.Error())
	}
	for i, datum := range pageData {
		var role model.Role
		_ = global.RoleColl.FindOne(context.TODO(), bson.M{"_id": datum["roleId"]}).Decode(&role)
		pageData[i]["role"] = role
		delete(pageData[i], "roleId")
	}
	//查询总数
	if total, err := global.UserColl.CountDocuments(context.TODO(), conditions); err != nil {
		return utils.ErrorMess("获取总数时失败", err.Error())
	} else {
		res := map[string]interface{}{
			"pageData": pageData,
			"total":    total,
		}
		return utils.SuccessMess("获取成功", res)
	}
}
