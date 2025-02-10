package router

import (
	"server/api"

	"github.com/gin-gonic/gin"
)

type FeedbackRouter struct {
}

func (f *FeedbackRouter) InitFeedbackRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup, AdminRouter *gin.RouterGroup) {
	feedbackRouter := Router.Group("feedback")
	feedbackPublicRouter := PublicRouter.Group("feedback")
	feedbackAdminRouter := AdminRouter.Group("feedback")

	feedbackApi := api.ApiGroupApp.FeedbackApi
	{
		feedbackRouter.POST("create", feedbackApi.FeedbackCreate)
		feedbackRouter.GET("info", feedbackApi.FeedbackInfo)
	}
	{
		feedbackPublicRouter.GET("new", feedbackApi.FeedbackNew)
	}
	{
		feedbackAdminRouter.DELETE("delete", feedbackApi.FeedbackDelete)
		feedbackAdminRouter.PUT("reply", feedbackApi.FeedbackReply)
		feedbackAdminRouter.GET("list", feedbackApi.FeedbackList)
	}
}
