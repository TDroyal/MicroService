package dao

import (
	"fmt"

	"github.com/redis/go-redis/v9"
	"gopkg.in/ini.v1"
)

// 操作redis

var RDB *redis.Client

func init() {
	config, iniErr := ini.Load("conf/app.ini")
	if iniErr != nil {
		fmt.Printf("Fail to read file: %v", iniErr)
		return
	}

	// 从ini文件中读出配置
	ip := config.Section("redis").Key("ip").String()
	port := config.Section("redis").Key("port").String()
	pw := config.Section("redis").Key("password").String()

	RDB = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", ip, port),
		Password: pw, // no password set
		DB:       0,  // use default DB
	})
}
