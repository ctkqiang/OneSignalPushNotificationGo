package main

import (
	"fmt"
	"valuefarm_pushnotification_services/internal/routes"
	"valuefarm_pushnotification_services/internal/utilities"

	"github.com/gin-gonic/gin"
)

var (
	Addr        = ":8080"
	Port        = 8080
)

func main() {
  gin.SetMode(gin.DebugMode)

	router := gin.Default()

	router.Use(gin.Recovery())
  router.Run(Addr)

  routes.StandardPushNotification(router, nil)

	if err := router.Run(fmt.Sprintf(":%d", Port)); err != nil {
		utilities.Log(utilities.ERROR, "HTTP 服务启动失败: %v", err)
	}
}