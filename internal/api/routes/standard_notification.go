package routes

import (
	"context"
	"net/http"
	"valuefarm_pushnotification_services/internal/api/handler"
	"valuefarm_pushnotification_services/internal/config"
	"valuefarm_pushnotification_services/internal/structure"
	"valuefarm_pushnotification_services/internal/utilities"

	"github.com/OneSignal/onesignal-go-api"
	"github.com/gin-gonic/gin"
)

// StandardPushNotification 注册标准推送通知路由
// @Summary 发送文本推送通知
// @Description 发送通用文本推送通知到所有用户
// @Tags 通知
// @Accept json
// @Produce json
// @Param notification body structure.NotificationContent true "通知内容"
// @Success 200 {object} map[string]interface{} "成功响应"
// @Failure 400 {object} map[string]interface{} "请求参数错误"
// @Failure 500 {object} map[string]interface{} "服务器内部错误"
// @Router /api/notification/text [post]
// @Example curl -X POST http://localhost:8080/api/notification/text \
// @Example   -H "Content-Type: application/json" \
// @Example   -d '{"title": "测试通知", "message": "这是一条测试通知"}'
func StandardPushNotification(router *gin.Engine) {
	public := router.Group(config.SEND_TEXT_PUSH_NOTIFICATION_HEAD)
	{
		public.POST(config.SEND_TEXT_PUSH_NOTIFICATION, SendTextPushNotification())
	}
}

func SendTextPushNotification() gin.HandlerFunc {
	return func(c *gin.Context) {
		var request structure.NotificationContent
		if err := c.ShouldBindJSON(&request); err != nil {
			utilities.Log(utilities.ERROR, "解析请求参数失败: %s", err.Error())
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  "error",
				"message": "Invalid request parameters",
			})
			return
		}

		client := &structure.OneSignalClient{
			APIClient:             onesignal.NewAPIClient(onesignal.NewConfiguration()),
			ApplicationId:         config.OneSignalCreds.AppID,
			AuthenticationContext: context.WithValue(context.Background(), onesignal.ContextAPIKeys, "Basic "+config.OneSignalCreds.APIKey),
		}

		response, err := handler.SendGeneralNotification(client, &request)
		if err != nil {
			utilities.Log(utilities.ERROR, "发送通知失败: %s", err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{
				"status":  "error",
				"message": "Failed to send notification",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"status":  "success",
			"message": "Notification sent successfully",
			"data":    response,
		})
	}
}