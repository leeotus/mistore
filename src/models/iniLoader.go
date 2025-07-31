package models

import (
	"fmt"

	"gopkg.in/ini.v1"
)

type IniLoader struct {
	RedisConfig RedisIni
	MysqlConfig MysqlIni
}

func (loader *IniLoader) LoadConfigs(path string) {
	cfg, err := ini.Load(path)
	if err != nil {
		panic(fmt.Sprintf("Failed to read .ini file: %v", path))
	}
	loader.MysqlConfig.HostAddr = cfg.Section("mysql").Key("host").String()
	loader.MysqlConfig.Username = cfg.Section("mysql").Key("username").String()
	loader.MysqlConfig.Password = cfg.Section("mysql").Key("password").String()
	loader.MysqlConfig.Port = cfg.Section("mysql").Key("port").String()
	loader.MysqlConfig.DBName = cfg.Section("mysql").Key("dbname").String()

	loader.RedisConfig.Host = cfg.Section("redis").Key("host").String()
	loader.RedisConfig.Username = cfg.Section("redis").Key("username").String()
	loader.RedisConfig.Password = cfg.Section("redis").Key("password").String()
	loader.RedisConfig.Port = cfg.Section("redis").Key("port").String()
}
