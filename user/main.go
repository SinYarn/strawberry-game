package main

import (
	"common/config"
	"common/metrics"
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"user/app"
)

// 启动命令
var configFile = flag.String("config", "application.yaml", "config file")

func main() {
	// 1. 加载配置
	flag.Parse()
	config.InitConfig(*configFile)
	fmt.Println(config.Conf)

	// 2. 启动性能监控 --内存
	go func() {
		err := metrics.Server(fmt.Sprintf("0.0.0.0:%d", config.Conf.MetricPort))
		if err != nil {
			panic(err)
		}
	}()

	// 3. 启动程序 grpc服务端
	err := app.Run(context.Background())
	if err != nil {
		log.Println(err)
		os.Exit(-1)
	}
}
