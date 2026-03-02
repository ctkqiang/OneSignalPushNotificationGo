package main

import (
	"flag"
	"fmt"
	"valuefarm_pushnotification_services/internal/api/routes"
	"valuefarm_pushnotification_services/internal/utilities"

	"github.com/gin-gonic/gin"
)

var (
	Addr        = ":8080"
	Port        = 8080
)

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
  router.Run(Addr)

  routes.StandardPushNotification(router, nil)

	if err := router.Run(fmt.Sprintf(":%d", Port)); err != nil {
		utilities.Log(utilities.ERROR, "HTTP 服务启动失败: %v", err)
	}
}