package router

import (
	"server/api"

	"github.com/gin-gonic/gin"
)

type AdvertisementRouter struct {
}

func (a *AdvertisementRouter) InitAdvertisementRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	advertisementRouter := Router.Group("advertisement")
	advertisementPublicRouter := PublicRouter.Group("advertisement")

	advertisementApi := api.ApiGroupApp.AdvertisementApi
	{
		advertisementRouter.POST("create", advertisementApi.AdvertisementCreate)
		advertisementRouter.DELETE("delete", advertisementApi.AdvertisementDelete)
		advertisementRouter.PUT("update", advertisementApi.AdvertisementUpdate)
		advertisementRouter.GET("list", advertisementApi.AdvertisementList)
	}
	{
		advertisementPublicRouter.GET("info", advertisementApi.AdvertisementInfo)
	}
}
