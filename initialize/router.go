package initialize

import (
	"net/http"
	"server/global"
	"server/middleware"
	"server/router"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

// InitRouter 初始化路由
func InitRouter() *gin.Engine {
	// 设置gin模式
	gin.SetMode(global.Config.System.Env)
	Router := gin.Default()
	// 使用日志记录中间件
	Router.Use(middleware.GinLogger(), middleware.GinRecovery(true))
	// 使用gin会话路由
	var store = cookie.NewStore([]byte(global.Config.System.SessionsSecret))
	Router.Use(sessions.Sessions("session", store))
	// 将指定目录下的文件提供给客户端
	// "uploads" 是URL路径前缀，http.Dir("uploads")是实际文件系统中存储文件的目录
	Router.StaticFS(global.Config.Upload.Path, http.Dir(global.Config.Upload.Path))
	// 创建路由组
	routerGroup := router.RouterGroupApp

	publicGroup := Router.Group(global.Config.System.RouterPrefix)
	privateGroup := Router.Group(global.Config.System.RouterPrefix)
	privateGroup.Use(middleware.JWTAuth())
	adminGroup := Router.Group(global.Config.System.RouterPrefix)
	adminGroup.Use(middleware.JWTAuth()).Use(middleware.AdminAuth())
	{
		routerGroup.InitBaseRouter(publicGroup)
	}
	{
		routerGroup.InitUserRouter(privateGroup, publicGroup, adminGroup)
		routerGroup.InitArticleRouter(privateGroup, publicGroup, adminGroup)
		routerGroup.InitCommentRouter(privateGroup, publicGroup, adminGroup)
		routerGroup.InitFeedbackRouter(privateGroup, publicGroup, adminGroup)
	}
	{
		routerGroup.InitImageRouter(adminGroup)
		routerGroup.InitAdvertisementRouter(adminGroup, publicGroup)
		routerGroup.InitFriendLinkRouter(adminGroup, publicGroup)
		routerGroup.InitWebsiteRouter(adminGroup, publicGroup)
		routerGroup.InitConfigRouter(adminGroup)
	}
	return Router
}
