package config

import (
	"fmt"
	"gopkg.in/ini.v1"
)

func init() {
	cfg, err := ini.Load("/config.ini")
	if err != nil {
		fmt.Println("无法读取配置文件,请检查配置文件:", err)
	}
	LadServe(cfg)
	LadData(cfg)
	LadRedis(cfg)
}

func LadServe(cfg *ini.File) {
	HttpHost = cfg.Section("HttpHost").Key("HttpHost").MustString("127.0.0.1")
	HttpPort = cfg.Section("HttpPort").Key("HttpPort").MustString("8081")
}

func LadData(cfg *ini.File) {
	DbType = cfg.Section("DbType").Key("DbType").MustString("mysql")
	DbHost = cfg.Section("DbHost").Key("DbHost").MustString("127.0.0.1")
	DbPort = cfg.Section("DbPort").Key("DbPort").MustString("3306")
	DbUSer = cfg.Section("DbUser").Key("DbUser").MustString("ximo")
	DbName = cfg.Section("DbName").Key("DBName").MustString("ximo")
}

func LadRedis(cfg *ini.File) {
	RdType = cfg.Section("RdType").Key("RdType").MustString("tcp")
	RdHost = cfg.Section("RdHost").Key("RdHost").MustString("127.0.0.1")
	RdPort = cfg.Section("RdPort").Key("RdPort").MustString("6379")
}
