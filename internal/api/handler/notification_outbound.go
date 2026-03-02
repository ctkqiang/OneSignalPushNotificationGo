package handler

import (
	"valuefarm_pushnotification_services/internal/structure"

	"github.com/OneSignal/onesignal-go-api"
)

func (s *structure.OneSignalClient) createBaseNotification(title, message string) *onesignal.Notification {
	notification := *onesignal.NewNotification(s.AppID)
	notification.SetContents(onesignal.StringMap{"en": &message})
	notification.SetHeadings(onesignal.StringMap{"en": &title})
	return &notification
}

func SendGeneralNotification() {}