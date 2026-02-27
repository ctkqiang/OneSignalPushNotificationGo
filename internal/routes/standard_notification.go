package routes

import (
	"valuefarm_pushnotification_services/internal/config"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)
func StandardPushNotification(router *gin.Engine, db *gorm.DB) {
	public := router.Group(config.SEND_TEXT_PUSH_NOTIFICATION_HEAD)
	{
		public.POST(config.SEND_TEXT_PUSH_NOTIFICATION, SendTextPushNotification(db))
	}
}

func SendTextPushNotification(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	}
}