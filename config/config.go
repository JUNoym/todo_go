package config

import (
	"log"
	"todo_app/utils"

	"gopkg.in/ini.v1"
)

type ConfigList struct {
	Port      string
	SQLDriver string
	DBName    string
	LogFile   string
}

var Config ConfigList

func init() {
	LoadConfig()
	utils.LoggingSettings(Config.LogFile)
}

func LoadConfig() {
	cfg, err := ini.Load("config/config.ini")
	if err != nil {
		log.Fatalln(err)
	}

	Config = ConfigList{
		Port:      cfg.Section("web").Key("port").MustString("8080"),
		SQLDriver: cfg.Section("db").Key("driver").MustString("sqlite3"),
		DBName:    cfg.Section("db").Key("name").MustString("webapp.sql"),
		LogFile:   cfg.Section("web").Key("logfile").MustString("webapp.log"),
	}
}
