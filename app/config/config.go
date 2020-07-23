package config

import "gopkg.in/ini.v1"

type ConfigList struct {
	Host 	 string
	Name 	 string
	User 	 string
	Password string
}

var Config ConfigList

func init()  {
	cfg, _ := ini.Load("config.ini")
	Config = ConfigList{
		Host: cfg.Section("db").Key("host").String(),
		Name: cfg.Section("db").Key("name").String(),
		User: cfg.Section("db").Key("user").String(),
		Password: cfg.Section("db").Key("password").String(),
	}
}