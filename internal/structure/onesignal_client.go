package structure

import (
	"context"
	"github.com/OneSignal/onesignal-go-api"
)

type OneSignalClient struct {
	APIClient *onesignal.APIClient
	ApplicationId     string
	AuthenticationContext   context.Context
}