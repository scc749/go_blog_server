package router

import (
	"server/api"
	"server/middleware"

	"github.com/gin-gonic/gin"
)

type UserRouter struct {
}

func (u *UserRouter) InitUserRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup, AdminRouter *gin.RouterGroup) {
	userRouter := Router.Group("user")
	userPublicRouter := PublicRouter.Group("user")
	userLoginRouter := PublicRouter.Group("user").Use(middleware.LoginRecord())
	userAdminRouter := AdminRouter.Group("user")
	userApi := api.ApiGroupApp.UserApi
	{
		userRouter.POST("logout", userApi.Logout)
		userRouter.PUT("resetPassword", userApi.UserResetPassword)
		userRouter.GET("info", userApi.UserInfo)
		userRouter.PUT("changeInfo", userApi.UserChangeInfo)
		userRouter.GET("weather", userApi.UserWeather)
		userRouter.GET("chart", userApi.UserChart)
	}
	{
		userPublicRouter.POST("forgotPassword", userApi.ForgotPassword)
		userPublicRouter.GET("card", userApi.UserCard)
	}
	{
		userLoginRouter.POST("register", userApi.Register)
		userLoginRouter.POST("login", userApi.Login)
	}
	{
		userAdminRouter.GET("list", userApi.UserList)
		userAdminRouter.PUT("freeze", userApi.UserFreeze)
		userAdminRouter.PUT("unfreeze", userApi.UserUnfreeze)
		userAdminRouter.GET("loginList", userApi.UserLoginList)
	}
}
