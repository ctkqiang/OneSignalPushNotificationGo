package main

import (
	"flag"
	"fmt"
	"pushnotification_services/internal/api/routes"
	"pushnotification_services/internal/utilities"

	_ "pushnotification_services/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

var (
	Addr        = ":8080"
	Port        = 8080
)

// @title 推送通知服务
// @version 1.0
// @description 提供推送通知相关的 API 接口
// @BasePath /
func main() {
	releaseMode := flag.Bool("release", false, "以发布模式运行服务器")
	flag.Parse()

	if *releaseMode {
		gin.SetMode(gin.ReleaseMode)
		fmt.Println("运行模式：RELEASE")
	} else {
		gin.SetMode(gin.DebugMode)
		fmt.Println("运行模式：DEBUG")
	}

	router := gin.Default()

	router.Use(gin.Recovery())

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	routes.InitWebSocketManager()
	routes.StandardPushNotification(router)
	routes.WebSocketRoutes(router)

	if err := router.Run(fmt.Sprintf(":%d", Port)); err != nil {
		utilities.Log(utilities.ERROR, "HTTP 服务启动失败: %v", err)
	}
}