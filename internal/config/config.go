package config

import (
	"os"
	"path/filepath"
	"valuefarm_pushnotification_services/internal/utilities"

	"github.com/joho/godotenv"
)

type LogMode string

var (
	LOG_DEBUG   = LogMode("DEBUG")
	LOG_VERBOSE = LogMode("VERBOSE")
	LOG_INFO    = LogMode("INFO")
)


type MySQLDatabase struct {
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
	MySQLCreds     MySQLDatabase
	OneSignalCreds OneSignalConfig
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

	MySQLCreds = MySQLDatabase{
		DatabaseUser:     os.Getenv("MYSQL_USER"),
		DatabasePassword: os.Getenv("MYSQL_PASSWORD"),
		DatabaseName:     os.Getenv("MYSQL_DATABASE"),
		DatabaseHost:     os.Getenv("MYSQL_HOST"),
		DatabasePort:     os.Getenv("MYSQL_PORT"),
	}

	OneSignalCreds = OneSignalConfig{
		AppID:  os.Getenv("ONESIGNAL_APP_ID"),
		APIKey: os.Getenv("ONESIGNAL_REST_API_KEY"),
	}
}
