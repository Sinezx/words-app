package util

import (
	"encoding/json"
	"os"
)

type Configuration struct {
	Dsn       string `json:"dsn"`
	AlarmLine int    `json:"alarmLine"`
}

var Config Configuration

func InitConfig() {
	AppPath, _ := os.Getwd()
	configFile, _ := os.Open(AppPath + ConfigFilePath)
	defer configFile.Close()
	decoder := json.NewDecoder(configFile)
	Config = Configuration{}
	decoder.Decode(&Config)
}
