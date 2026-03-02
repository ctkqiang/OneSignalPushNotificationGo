package handler

import (
	"pushnotification_services/internal/database"
	"pushnotification_services/internal/repositories"
	"pushnotification_services/internal/structure"
	"pushnotification_services/internal/utilities"

	"github.com/OneSignal/onesignal-go-api"
)

func CreateBaseNotification(appID, title, message string) *onesignal.Notification {
	notification := *onesignal.NewNotification(appID)

	contents := onesignal.StringMap{}
	contents.SetEn(message)
	
	headings := onesignal.StringMap{}
	headings.SetEn(title)
	
	notification.SetContents(contents)
	notification.SetHeadings(headings)
	
	return &notification
}



// SendGeneralNotification 发送通用通知
// 此函数仅用于发送通用通知，只包含标题和消息文本，不包含图片
// 通知将广播给所有用户，不针对特定分段
// 参数:
//   o: OneSignalClient 实例，用于与 OneSignal API 交互
//   s: NotificationContent 实例，包含通知的标题和消息内容
// 返回值:
//   database.NotificationResponse: 通知发送的响应信息，包含状态和内容
//   error: 发送过程中遇到的错误，如果成功则为 nil
func SendGeneralNotification(client *structure.OneSignalClient, content *structure.NotificationContent) (database.NotificationResponse, error) {
	utilities.Log(utilities.INFO, "正在发送通用通知")

	notification := *onesignal.NewNotification(client.ApplicationId)

	contents := onesignal.StringMap{}
	contents.SetEn(content.Message)

	headings := onesignal.StringMap{}
	headings.SetEn(content.Title)

	notification.SetContents(contents)
	notification.SetHeadings(headings)

	allUsersSegment := []string{"All"}
	notification.SetIncludedSegments(allUsersSegment)

	_, _, err := client.APIClient.DefaultApi.CreateNotification(client.AuthenticationContext).Notification(notification).Execute()
	if err != nil {
		utilities.Log(utilities.ERROR, "%s", "发送通用通知失败: "+err.Error())
		return database.NotificationResponse{
			Status: database.StatusFailed,
		}, 
		err
	}

	utilities.Log(utilities.INFO, "通用通知发送成功")

	if err := repositories.SaveRecord(&database.NotificationResponse{
		Status: database.StatusSuccess,
		Content: content,
	}); err != nil {
		utilities.Log(utilities.ERROR, "保存通知记录失败: %s", err.Error())
		return database.NotificationResponse{
			Status: database.StatusFailed,
		}, 
		err
	}
	
	return database.NotificationResponse{
		Status: database.StatusSuccess,
		Content: content,
	}, nil
}
