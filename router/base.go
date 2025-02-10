package router

import (
	"server/api"

	"github.com/gin-gonic/gin"
)

type BaseRouter struct {
}

func (b *BaseRouter) InitBaseRouter(Router *gin.RouterGroup) {
	baseRouter := Router.Group("base")

	baseApi := api.ApiGroupApp.BaseApi
	{
		baseRouter.POST("captcha", baseApi.Captcha)
		baseRouter.POST("sendEmailVerificationCode", baseApi.SendEmailVerificationCode)
		baseRouter.GET("qqLoginURL", baseApi.QQLoginURL)
	}
}
