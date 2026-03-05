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
		public.GET(config.INDEX, Index())
		public.GET(config.ENV, Env())
		public.GET(config.HEALTH, CurrentHealth())
	}
}

func Env() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status":  "ok",
			"service": "push-notification-service",
			"message": "哎哟喂, 你真以为我会把 .env 放在这儿让你黑掉这服务？",
			"timestamp": gin.H{
				"unix":  time.Now().Unix(),
				"utc":   time.Now().UTC(),
			},
		})
	}
}

func Index() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status":  "ok",
			"service": "push-notification-service",
			"message": "哟，您这是想推个通知？别墨迹，麻溜儿的！",
			"timestamp": gin.H{
				"unix":  time.Now().Unix(),
				"utc":   time.Now().UTC(),
			},
		})
	}
}

func CurrentHealth() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status":  "healthy",
			"service": "push-notification-service",
			"timestamp": gin.H{
				"unix":  time.Now().Unix(),
				"human": time.Now().Format(time.RFC3339),
			},
		})
	}
}