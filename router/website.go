package router

import (
	"server/api"

	"github.com/gin-gonic/gin"
)

type WebsiteRouter struct {
}

func (w *WebsiteRouter) InitWebsiteRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	websiteRouter := Router.Group("website")
	websitePublicRouter := PublicRouter.Group("website")

	websiteApi := api.ApiGroupApp.WebsiteApi
	{
		websiteRouter.POST("addCarousel", websiteApi.WebsiteAddCarousel)
		websiteRouter.PUT("cancelCarousel", websiteApi.WebsiteCancelCarousel)
		websiteRouter.POST("createFooterLink", websiteApi.WebsiteCreateFooterLink)
		websiteRouter.DELETE("deleteFooterLink", websiteApi.WebsiteDeleteFooterLink)
	}
	{
		websitePublicRouter.GET("logo", websiteApi.WebsiteLogo)
		websitePublicRouter.GET("title", websiteApi.WebsiteTitle)
		websitePublicRouter.GET("info", websiteApi.WebsiteInfo)
		websitePublicRouter.GET("carousel", websiteApi.WebsiteCarousel)
		websitePublicRouter.GET("news", websiteApi.WebsiteNews)
		websitePublicRouter.GET("calendar", websiteApi.WebsiteCalendar)
		websitePublicRouter.GET("footerLink", websiteApi.WebsiteFooterLink)
	}
}
