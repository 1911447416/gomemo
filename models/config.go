package models

import (
	"fmt"

	"gopkg.in/ini.v1"
)

type AppConfig struct {
	Port        string `ini:"port"`
	MysqlConfig `ini:"mysql"`
}

type MysqlConfig struct {
	Ip       string `ini:"ip"`
	Port     string `ini:"port"`
	User     string `ini:"user"`
	Password string `ini:"password"`
	Database string `ini:"database"`
}

func LoadConfig() *AppConfig {
	// 读取配置文件
	cfgfile := LoadConfDocument()
	// map至结构体
	var appConfig = new(AppConfig)
	err = cfgfile.MapTo(appConfig)
	if err != nil {
		fmt.Println("失败")
	}
	return appConfig
}

func LoadConfDocument() *ini.File {
	// 读取配置文件
	cfgfile, err := ini.Load("./conf/app.ini")
	if err != nil {
		fmt.Println("fail to load conf.err:", err)
	}
	return cfgfile
}
