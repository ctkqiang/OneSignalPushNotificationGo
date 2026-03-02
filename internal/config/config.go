package config

import (
	"os"
	"path/filepath"
	"pushnotification_services/internal/utilities"

	"github.com/joho/godotenv"
)

type LogMode string

var (
	LOG_DEBUG   = LogMode("DEBUG")
	LOG_VERBOSE = LogMode("VERBOSE")
	LOG_INFO    = LogMode("INFO")
)


type MongoDBConfig struct {
	DatabaseUser     string `json:"database_user"`
	DatabasePassword string `json:"database_password"`
	DatabaseName     string `json:"database_name"`
	DatabaseHost     string `json:"database_host"`
	DatabasePort     string `json:"database_port"`
}


type OneSignalConfig struct {
	AppID  string
	APIKey string
}

var (
	COLLECTION_NOTIFICATIONS = "notifications_log"
)

var (
	MongoDBCreds    MongoDBConfig
	OneSignalCreds  OneSignalConfig
)

func init() {
	selectedEnviornmentFlavours, err:=  GetDevelopmentFlavours()
	if err != nil {
		utilities.Log(utilities.ERROR, "Environment Error %s", err.Error())
	}

	workspaceDirectory, err := os.Getwd()
	
	if err == nil {
		envPath := filepath.Join(workspaceDirectory, "internal", "config", selectedEnviornmentFlavours)

		if _, err := os.Stat(envPath); err == nil {
			if err := godotenv.Load(envPath); err != nil {
				utilities.Log(utilities.ERROR, "[CONFIG] WARNING!!! Failed to load [.env] File %s: %v\n", envPath, err)
			}
		} else {
			godotenv.Load()
		}
	}

	MongoDBCreds = MongoDBConfig{
		DatabaseUser:     os.Getenv("MONGODB_USER"),
		DatabasePassword: os.Getenv("MONGODB_PASSWORD"),
		DatabaseName:     os.Getenv("MONGODB_DATABASE"),
		DatabaseHost:     os.Getenv("MONGODB_HOST"),
		DatabasePort:     os.Getenv("MONGODB_PORT"),
	}

	OneSignalCreds = OneSignalConfig{
		AppID:  os.Getenv("ONESIGNAL_APP_ID"),
		APIKey: os.Getenv("ONESIGNAL_REST_API_KEY"),
	}
}
