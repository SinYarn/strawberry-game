package app

import (
	"common/config"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// Run 启动程序 启动grpc服务 启动http服务 启动日志服务 启动数据库
func Run(ctx context.Context) error {
	// 1.Todo: 日志库 info error fatal debug

	// 2.Todo: etcd 注册中心 grpc注册到etcd 客户端访问的时候通过etcd(负载均衡) 获取grpc的地址, 分布式应用 微服务的基础使用

	// 3. 启用grpc服务端
	server := grpc.NewServer()
	lis, err := net.Listen("tcp", config.Conf.Grpc.Addr)
	if err != nil {
		log.Fatalf("user grpc server failed to listen: %v", err)
	}
	// Todo: 注册grpc service 需要数据库  mongo redis
	// Todo初始化 数据库管理
	// 阻塞
	go func() {
		err = server.Serve(lis)
		if err != nil {
			log.Fatalf("user grpc server run failed err:%v", err)
		}
	}()

	// 优雅启停 遇到中断信号 --> 退出
	stop := func() {
		server.Stop()

		// other 停止3s
		time.Sleep(3 * time.Second)
		fmt.Println("user grpc server stop")
	}

	// 信号量
	// os.Signal 是 Go 标准库中定义的一个类型，用于表示操作系统发出的信号，例如终止信号（SIGTERM）、中断信号（SIGINT）等。
	// 通过创建这个 channel，程序可以接收这些信号并做出相应的处理。
	c := make(chan os.Signal, 1)
	// 监听信号 中断 退出 终止 挂断
	signal.Notify(c, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGINT, syscall.SIGHUP)

	for {
		select {
		// 上下文被取消
		case <-ctx.Done():
			stop()
			// time out
			return nil
		case s := <-c:
			switch s {
			case syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGINT:
				stop()
				log.Println("user app quit!")
				return nil
			// session 用户退出触发
			case syscall.SIGHUP:
				stop()
				log.Println("user app hang up!, user app quit")
				return nil
			default:
				return nil
			}
		}
	}
	return nil
}
