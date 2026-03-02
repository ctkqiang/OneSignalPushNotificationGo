package routes

import (
	"pushnotification_services/internal/api/handler"
	"pushnotification_services/internal/config"
	"pushnotification_services/internal/structure"

	"github.com/gin-gonic/gin"
)

var WebSocketManager *structure.WebSocketManager

func InitWebSocketManager() {
	WebSocketManager = handler.NewWebSocketManager()
	go handler.Run(WebSocketManager)
}

func WebSocketRoutes(router *gin.Engine) {
	if WebSocketManager == nil {
		InitWebSocketManager()
	}

	router.GET(config.WEBSCOKET_CHANNEL, func(c *gin.Context) {
		handler.HandleWebSocket(WebSocketManager, c)
	})
}
