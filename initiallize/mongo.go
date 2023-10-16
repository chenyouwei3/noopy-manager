package initiallize

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"noopy-manager/global"
)

func MongoInit() {
	if global.MongoClientYD == nil {
		global.MongoClientYD = getMongoClient("")
	}
}

func getMongoClient(uri string) *mongo.Client {
	clientOptions := options.Client().ApplyURI(uri)
	// 使用客户端选项和上下文（context.TODO()）连接到 MongoDB 数据库
	MongoClient, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Println(err)
		fmt.Println("g")
	}
	// 使用 Ping 方法检测与数据库的连接状态
	if err = MongoClient.Ping(context.TODO(), nil); err != nil {
		log.Println(err)
		fmt.Println("g")
	}
	fmt.Println("mongodb连接成功")
	return MongoClient
}
