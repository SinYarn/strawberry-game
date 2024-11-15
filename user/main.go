package main

import (
	"common/config"
	"flag"
)

// 启动命令
var configFile = flag.String("config", "application.yaml", "config file")

func main() {
	// 1. 加载配置
	flag.Parse()
	config.InitConfig(*configFile)
	// 2. 启动监控 --内存
	// 3. 启动程序 grpc服务端
}
