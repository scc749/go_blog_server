package api

import "server/service"

type ApiGroup struct {
	BaseApi
	UserApi
	ImageApi
	ArticleApi
	CommentApi
	AdvertisementApi
	FriendLinkApi
	FeedbackApi
	WebsiteApi
	ConfigApi
}

var ApiGroupApp = new(ApiGroup)

var baseService = service.ServiceGroupApp.BaseService
var userService = service.ServiceGroupApp.UserService
var qqService = service.ServiceGroupApp.QQService
var jwtService = service.ServiceGroupApp.JwtService
var imageService = service.ServiceGroupApp.ImageService
var articleService = service.ServiceGroupApp.ArticleService
var commentService = service.ServiceGroupApp.CommentService
var advertisementService = service.ServiceGroupApp.AdvertisementService
var friendLinkService = service.ServiceGroupApp.FriendLinkService
var feedbackService = service.ServiceGroupApp.FeedbackService
var websiteService = service.ServiceGroupApp.WebsiteService
var configService = service.ServiceGroupApp.ConfigService
