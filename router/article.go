package router

import (
	"server/api"

	"github.com/gin-gonic/gin"
)

type ArticleRouter struct {
}

func (a *ArticleRouter) InitArticleRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup, AdminRouter *gin.RouterGroup) {
	articleRouter := Router.Group("article")
	articlePublicRouter := PublicRouter.Group("article")
	articleAdminRouter := AdminRouter.Group("article")

	articleApi := api.ApiGroupApp.ArticleApi
	{
		articleRouter.POST("like", articleApi.ArticleLike)
		articleRouter.GET("isLike", articleApi.ArticleIsLike)
		articleRouter.GET("likesList", articleApi.ArticleLikesList)
	}
	{
		articlePublicRouter.GET(":id", articleApi.ArticleInfoByID)
		articlePublicRouter.GET("search", articleApi.ArticleSearch)
		articlePublicRouter.GET("category", articleApi.ArticleCategory)
		articlePublicRouter.GET("tags", articleApi.ArticleTags)
	}
	{
		articleAdminRouter.POST("create", articleApi.ArticleCreate)
		articleAdminRouter.DELETE("delete", articleApi.ArticleDelete)
		articleAdminRouter.PUT("update", articleApi.ArticleUpdate)
		articleAdminRouter.GET("list", articleApi.ArticleList)
	}
}
