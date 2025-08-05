package models

import (
	"fmt"

	"gopkg.in/ini.v1"
)

type IniLoader struct {
	RedisConfig  RedisIni
	MysqlConfig  MysqlIni
	AliOOSConfig AliOOSIni
}

var Loader = &IniLoader{}

func LoadConfigs(path string) {
	cfg, err := ini.Load(path)
	if err != nil {
		panic(fmt.Sprintf("Failed to read .ini file: %v", path))
	}
	Loader.MysqlConfig.HostAddr = cfg.Section("mysql").Key("host").String()
	Loader.MysqlConfig.Username = cfg.Section("mysql").Key("username").String()
	Loader.MysqlConfig.Password = cfg.Section("mysql").Key("password").String()
	Loader.MysqlConfig.Port = cfg.Section("mysql").Key("port").String()
	Loader.MysqlConfig.DBName = cfg.Section("mysql").Key("dbname").String()

	Loader.RedisConfig.Host = cfg.Section("redis").Key("host").String()
	Loader.RedisConfig.Username = cfg.Section("redis").Key("username").String()
	Loader.RedisConfig.Password = cfg.Section("redis").Key("password").String()
	Loader.RedisConfig.Port = cfg.Section("redis").Key("port").String()

	Loader.AliOOSConfig.EndPoint = cfg.Section("oss").Key("endpoint").String()
	Loader.AliOOSConfig.AccessKey = cfg.Section("oss").Key("accessKey").String()
	Loader.AliOOSConfig.AccessSecret = cfg.Section("oss").Key("accessSecret").String()
	Loader.AliOOSConfig.Domain = cfg.Section("oss").Key("domain").String()
}
