package config

import (
	"log"
	"os"

	ini "gopkg.in/ini.v1"
)

// ConfigList is api key struct
type ConfigList struct {
	FootballDataApiToken string
	FootballDataBaseUrl  string
}

// Config is ConfigList
var Config ConfigList

func init() {
	cfg, err := ini.Load("config/config.ini")
	if err != nil {
		log.Printf("Failed to read file: %v", err)
		os.Exit(1)
	}

	Config = ConfigList{
		FootballDataApiToken: cfg.Section("footballData").Key("apiToken").String(),
		FootballDataBaseUrl:  cfg.Section("footballData").Key("baseUrl").String(),
	}
}
