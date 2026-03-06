package service

import (
	"context"
	"fmt"
	"pushnotification_services/internal/config"
	"pushnotification_services/internal/utilities"
	"sync"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	mongoClient *mongo.Client
	mongoOnce   sync.Once
	mongoErr    error
)

// GetMongoDatabaseConnection 获取 MongoDB 数据库连接
// 返回值:
//   *mongo.Client: MongoDB 客户端实例
//   error: 连接过程中遇到的错误，如果成功则为 nil
func GetMongoDatabaseConnection() (*mongo.Client, error) {
	mongoOnce.Do(func() {
		creds := config.MongoDBCreds

		var connectionString string
		if creds.DatabaseUser != "" && creds.DatabasePassword != "" {
			connectionString = fmt.Sprintf("mongodb://%s:%s@%s:%s/%s",
				creds.DatabaseUser,
				creds.DatabasePassword,
				creds.DatabaseHost,
				creds.DatabasePort,
				creds.DatabaseName,
			)
		} else {
			connectionString = fmt.Sprintf("mongodb://%s:%s/%s",
				creds.DatabaseHost,
				creds.DatabasePort,
				creds.DatabaseName,
			)
		}

		clientOptions := options.Client().ApplyURI(connectionString)

		client, err := mongo.Connect(context.Background(), clientOptions)
		if err != nil {
			utilities.Log(utilities.ERROR, "MongoDB 连接失败: %s", err.Error())
			mongoErr = err
			return
		}

		err = client.Ping(context.Background(), nil)
		if err != nil {
			utilities.Log(utilities.ERROR, "MongoDB  ping 失败: %s", err.Error())
			mongoErr = err
			return
		}

		utilities.Log(utilities.INFO, "MongoDB 连接成功")
		mongoClient = client
	})

	return mongoClient, mongoErr
}

// GetMongoClient 获取已存在的 MongoDB 客户端实例（如果有）
// 返回值:
//   *mongo.Client: MongoDB 客户端实例，如果未初始化则为 nil
func GetMongoClient() *mongo.Client {
	return mongoClient
}