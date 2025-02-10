//go:build !windows
// +build !windows

package core

import (
	"time"

	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
)

// initServer 函数初始化一个 Endless 服务器（适用于非 Windows 系统）
func initServer(address string, router *gin.Engine) server {
	s := endless.NewServer(address, router) // 使用 endless 包创建一个新的 HTTP 服务器实例
	s.ReadHeaderTimeout = 10 * time.Minute  // 设置请求头的读取超时时间为 10 分钟
	s.WriteTimeout = 10 * time.Minute       // 设置响应写入的超时时间为 10 分钟
	s.MaxHeaderBytes = 1 << 20              // 设置最大请求头的大小（1MB）

	return s // 返回创建的服务器实例
}
