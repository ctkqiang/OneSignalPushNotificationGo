package routes

import (
	"net/http"
	"pushnotification_services/internal/config"
	"time"

	"github.com/gin-gonic/gin"
)

func General(router *gin.Engine) {
	public := router.Group(config.INDEX)
	{
		public.GET(config.HEALTH, CurrentHealth())
	}
}

func CurrentHealth() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "healthy",
			"service": "push-notification-service",
			"timestamp": time.Now().Format(time.RFC3339),
		})
	}
}