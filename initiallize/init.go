package initiallize

import (
	"context"
	"fmt"
	"noopy-manager/global"
)

func Init() {
	MongoInit()
	RedisInit()
}

func DataBaseClose() {
	err := global.RedisClient.Close()
	if err != nil {
		fmt.Println("Error on closing redisService client.")
	}
	err = global.MongoClientYD.Disconnect(context.TODO())
	if err != nil {
		fmt.Println("Error on closing MongoDB connection:")
	}
}
