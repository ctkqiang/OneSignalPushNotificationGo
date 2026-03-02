package service

import (
	"context"
	"pushnotification_services/internal/config"
	"pushnotification_services/internal/structure"
	"pushnotification_services/internal/utilities"

	"github.com/OneSignal/onesignal-go-api"
)
func OneSignalConnection() *structure.OneSignalClient{ 
	var (
		appID = config.OneSignalCreds.AppID
		apiKey = config.OneSignalCreds.APIKey
	)

	if appID == "" {
		utilities.Log(utilities.ERROR, "环境变量中缺少 OneSignal AppID")
	}

	if apiKey == "" {
		utilities.Log(utilities.ERROR, "环境变量中缺少 OneSignal APIKey")
	}

	configuration := onesignal.NewConfiguration()
	client := onesignal.NewAPIClient(configuration)

	authenticationContext := context.WithValue(
		context.Background(),
		onesignal.AppAuth,
		apiKey,
	)

	return &structure.OneSignalClient{
		APIClient: client,
		ApplicationId: appID,
		AuthenticationContext: authenticationContext,
	}
}