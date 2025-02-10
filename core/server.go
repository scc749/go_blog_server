package core

import (
	"server/global"
	"server/initialize"
	"server/service"

	"go.uber.org/zap"
)

type server interface {
	ListenAndServe() error
}

// RunServer 用于启动服务器
func RunServer() {
	addr := global.Config.System.Addr()
	Router := initialize.InitRouter()

	// 加载所有的 JWT 黑名单，存入本地缓存
	service.LoadAll()

	// 初始化服务器并启动
	s := initServer(addr, Router)
	global.Log.Info("server run success on ", zap.String("address", addr))
	global.Log.Error(s.ListenAndServe().Error())
}
