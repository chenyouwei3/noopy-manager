package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Api struct {
	Id     primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Name   string             `bson:"name" json:"name"`     //api名称
	Url    string             `bson:"url" json:"url"`       //路由
	Method string             `bson:"method" json:"method"` //方法
	Desc   string             `bson:"desc" json:"desc"`     //描述
}
