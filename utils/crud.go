package utils

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

// Find 查询
func Find(collection *mongo.Collection, result interface{}, filter interface{}, opts ...*options.FindOptions) error {
	cur, err := collection.Find(context.TODO(), filter, opts...)
	if err != nil {
		return err
	}
	if err = cur.All(context.TODO(), result); err != nil {
		return err
	}
	return nil
}
func FindOne(collection *mongo.Collection, filter interface{}, opts ...*options.FindOneOptions) Response {
	var result bson.M
	err := collection.FindOne(context.TODO(), filter, opts...).Decode(&result)
	if err != nil {
		return ErrorMess("查找失败", err.Error())
	} else {
		return SuccessMess("查找成功", result)
	}
}
func InsertOne(collection *mongo.Collection, document interface{}, opts ...*options.InsertOneOptions) Response {
	if res, err := collection.InsertOne(context.TODO(), document, opts...); err != nil {
		return ErrorMess("添加失败", err.Error())
	} else {
		return SuccessMess("添加成功", res)
	}
}
func InsertMany(collection *mongo.Collection, document []interface{}, opts ...*options.InsertManyOptions) Response {
	if res, err := collection.InsertMany(context.TODO(), document, opts...); err != nil {
		return ErrorMess("添加失败", err.Error())
	} else {
		return SuccessMess("添加成功", res)
	}
}
func DeleteOne(collection *mongo.Collection, filter interface{}, opts ...*options.DeleteOptions) Response {
	res, err := collection.DeleteOne(context.TODO(), filter, opts...)
	if err != nil {
		return ErrorMess("删除失败", err.Error())
	}
	if res.DeletedCount == 0 {
		return ErrorMess("删除不存在", res)
	}
	return SuccessMess("删除成功", res)
}
func UpdateOne(collection *mongo.Collection, filter interface{}, update interface{}, opts ...*options.FindOneAndUpdateOptions) Response {
	if err := collection.FindOneAndUpdate(context.TODO(), filter, update, opts...).Decode(&bson.M{}); err != nil {
		return ErrorMess("更新失败", err.Error())
	} else {
		return SuccessMess("更新成功", filter)
	}
}
func Aggregate(collection *mongo.Collection, result interface{}, filter interface{}, opts ...*options.AggregateOptions) error {
	cur, err := collection.Aggregate(context.TODO(), filter, opts...)
	if err != nil {
		return err
	}
	if err = cur.All(context.TODO(), result); err != nil {
		return err
	}
	return nil
}
func GetPageData(collection *mongo.Collection, conditions map[string]interface{}, pageSize, currPage int64) Response {
	var pageData []map[string]interface{}
	skip := (currPage - 1) * pageSize
	//获取分页数据
	if err := Find(collection, &pageData, conditions, &options.FindOptions{
		Limit: &pageSize,
		Skip:  &skip,
		Sort:  bson.M{"_id": -1},
	}); err != nil {
		return ErrorMess("获取分页数据失败", err.Error())
	}
	//查询总数
	if total, err := collection.CountDocuments(context.TODO(), conditions); err != nil {
		return ErrorMess("获取总数时失败", err.Error())
	} else {
		res := map[string]interface{}{
			"pageData": pageData,
			"total":    total,
		}
		return SuccessMess("获取成功", res)
	}
}

// 得到数据总条数
func GetCount(col *mongo.Collection, searchKey, value string) int64 {
	totalCount, err := col.CountDocuments(context.TODO(), bson.M{searchKey: primitive.Regex{Pattern: value}})
	if err != nil {
		log.Println(err)
	}
	return totalCount
}

// 得到数据总条数
func GetCountByTime(col *mongo.Collection, startTime, endTime string) int64 {
	if startTime == "" || endTime == "" {
		totalCount, err := col.CountDocuments(context.TODO(), bson.M{})
		if err != nil {
			log.Println(err)
		}
		return totalCount
	} else {
		totalCount, err := col.CountDocuments(context.TODO(), bson.M{"createTime": bson.M{"$gte": startTime, "$lte": endTime}})
		if err != nil {
			log.Println(err)
		}
		return totalCount
	}
}
