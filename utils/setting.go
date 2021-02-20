package utils

import (
	"log"

	ini "gopkg.in/ini.v1"
)

var (
	Cfg *ini.File

	AppMode  string
	HttpPort string

	Db         string
	Dbhost     string
	DbPort     string
	DbUser     string
	DbPassword string
	DbName     string
)

func init() {
	var err error
	Cfg, err = ini.Load("config/config.ini")
	if err != nil {
		log.Fatalf("Fail to parse 'config/config.ini': %v", err)
	}
	LoadServer()
	LoadData()
}

func LoadServer() {
	AppMode = Cfg.Section("server").Key("AppMode").MustString("debug")
	HttpPort = Cfg.Section("server").Key("HttpPort").MustString("3000")
}

func LoadData() {
	Db = Cfg.Section("database").Key("Db").MustString("mysql")
	Dbhost = Cfg.Section("database").Key("Dbhost").MustString("localhost")
	DbPort = Cfg.Section("database").Key("DbPort").MustString("3306")
	DbUser = Cfg.Section("database").Key("DbUser").MustString("root")
	DbPassword = Cfg.Section("database").Key("DbPassword").MustString("123456")
	DbName = Cfg.Section("database").Key("DbName").MustString("blog")
}
