package router

import (
	"server/api"

	"github.com/gin-gonic/gin"
)

type FriendLinkRouter struct {
}

func (f *FriendLinkRouter) InitFriendLinkRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	friendLinkRouter := Router.Group("friendLink")
	friendLinkPublicRouter := PublicRouter.Group("friendLink")

	friendLinkApi := api.ApiGroupApp.FriendLinkApi
	{
		friendLinkRouter.POST("create", friendLinkApi.FriendLinkCreate)
		friendLinkRouter.DELETE("delete", friendLinkApi.FriendLinkDelete)
		friendLinkRouter.PUT("update", friendLinkApi.FriendLinkUpdate)
		friendLinkRouter.GET("list", friendLinkApi.FriendLinkList)
	}
	{
		friendLinkPublicRouter.GET("info", friendLinkApi.FriendLinkInfo)
	}
}
