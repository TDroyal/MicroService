package main

import (
	"captcha/handler"
	pb "captcha/proto/captcha"
	"fmt"

	"github.com/go-micro/plugins/v4/registry/consul"
	"go-micro.dev/v4"
	"go-micro.dev/v4/logger"
	"go-micro.dev/v4/registry"
	"gopkg.in/ini.v1"
)

var (
	service = "captcha"
	version = "latest"
)

func main() {
	// 集成consul服务发现
	config, iniErr := ini.Load("conf/app.ini")
	if iniErr != nil {
		fmt.Printf("Fail to read file: %v", iniErr)
		return
	}

	// 从ini文件中读出配置
	ip := config.Section("consul").Key("ip").String()
	port := config.Section("consul").Key("port").String()

	consulRegistry := consul.NewRegistry(
		registry.Addrs(fmt.Sprintf("%s:%s", ip, port)),
	)

	// Create service
	srv := micro.NewService()
	srv.Init(
		micro.Name(service),
		micro.Version(version),
		micro.Registry(consulRegistry), //自己配置consul
	)

	// Register handler
	if err := pb.RegisterCaptchaHandler(srv.Server(), new(handler.Captcha)); err != nil {
		logger.Fatal(err)
	}
	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
