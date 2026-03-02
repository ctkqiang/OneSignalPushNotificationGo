package service

import (
	"context"
	"fmt"
	"pushnotification_services/internal/config"
	"pushnotification_services/internal/utilities"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// GetMongoDatabaseConnection 建立 MongoDB 数据库连接
// 返回值:
//   *mongo.Client: MongoDB 客户端实例
//   error: 连接过程中遇到的错误，如果成功则为 nil
func GetMongoDatabaseConnection() (*mongo.Client, error) {
	creds := config.MongoDBCreds

	connectionString := fmt.Sprintf("mongodb://%s:%s@%s:%s/%s",
		creds.DatabaseUser,
		creds.DatabasePassword,
		creds.DatabaseHost,
		creds.DatabasePort,
		creds.DatabaseName,
	)

	clientOptions := options.Client().ApplyURI(connectionString)

	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		utilities.Log(utilities.ERROR, "MongoDB 连接失败: %s", err.Error())
		return nil, err
	}

	err = client.Ping(context.Background(), nil)
	if err != nil {
		utilities.Log(utilities.ERROR, "MongoDB  ping 失败: %s", err.Error())
		return nil, err
	}

	utilities.Log(utilities.INFO, "MongoDB 连接成功")
	return client, nil
}