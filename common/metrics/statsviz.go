package metrics

import (
	"github.com/arl/statsviz"
	"net/http"
)

// Server 可视化实时监控
// http://host:5854/debug/statsviz/
func Server(add string) error {
	// 创建一个新的 HTTP 多路复用器（ServeMux）
	mux := http.NewServeMux()
	// 实时监控 Go 应用程序的库, 提供了一些可视化工具来监控内存使用、CPU 使用等性能指标。
	if err := statsviz.Register(mux); err != nil {
		return err
	}
	// 启动一个 HTTP 服务器，监听指定的地址 add，并使用 mux 作为请求处理器。
	if err := http.ListenAndServe(add, mux); err != nil {
		return err
	}
	return nil
}
