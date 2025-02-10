package api

import (
	"net/http"
	"server/global"
	"server/model/database"
	"server/model/request"
	"server/model/response"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type WebsiteApi struct {
}

// WebsiteLogo 网站 Logo 链接
func (website *WebsiteApi) WebsiteLogo(c *gin.Context) {
	if global.Config.Website.Logo != "" {
		c.Redirect(http.StatusMovedPermanently, global.Config.Website.Logo)
	} else {
		c.Redirect(http.StatusMovedPermanently, "/image/logo.png")
	}
}

// WebsiteTitle 网站标题栏
func (website *WebsiteApi) WebsiteTitle(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"title": global.Config.Website.Title})
}

// WebsiteInfo 获取网站信息
func (website *WebsiteApi) WebsiteInfo(c *gin.Context) {
	response.OkWithData(global.Config.Website, c)
}

// WebsiteCarousel 获取首页背景
func (website *WebsiteApi) WebsiteCarousel(c *gin.Context) {
	urls := websiteService.WebsiteCarousel()
	response.OkWithData(urls, c)
}

// WebsiteNews 获取新闻
func (website *WebsiteApi) WebsiteNews(c *gin.Context) {
	sourceStr := c.Query("source")
	hotSearchData, err := websiteService.WebsiteNews(sourceStr)
	if err != nil {
		global.Log.Error("Failed to get news:", zap.Error(err))
		response.FailWithMessage("Failed to get news", c)
		return
	}
	response.OkWithData(hotSearchData, c)
}

// WebsiteCalendar 获取日历
func (website *WebsiteApi) WebsiteCalendar(c *gin.Context) {
	dateStr := time.Now().Format("2006/0102")
	calendar, err := websiteService.WebsiteCalendar(dateStr)
	if err != nil {
		global.Log.Error("Failed to get calendar:", zap.Error(err))
		response.FailWithMessage("Failed to get calendar", c)
		return
	}
	response.OkWithData(calendar, c)
}

// WebsiteFooterLink 获取页脚链接
func (website *WebsiteApi) WebsiteFooterLink(c *gin.Context) {
	footerLinks := websiteService.WebsiteFooterLink()
	response.OkWithData(footerLinks, c)
}

// WebsiteAddCarousel 添加首页背景
func (website *WebsiteApi) WebsiteAddCarousel(c *gin.Context) {
	var req request.WebsiteCarouselOperation
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	err = websiteService.WebsiteAddCarousel(req)

	if err != nil {
		global.Log.Error("Failed to add carousel:", zap.Error(err))
		response.FailWithMessage("Failed to add carousel", c)
		return
	}
	response.OkWithMessage("Successfully added carousel", c)
}

// WebsiteCancelCarousel 移除首页背景
func (website *WebsiteApi) WebsiteCancelCarousel(c *gin.Context) {
	var req request.WebsiteCarouselOperation
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	err = websiteService.WebsiteCancelCarousel(req)

	if err != nil {
		global.Log.Error("Failed to cancel carousel:", zap.Error(err))
		response.FailWithMessage("Failed to cancel carousel", c)
		return
	}
	response.OkWithMessage("Successfully canceled carousel", c)
}

// WebsiteCreateFooterLink 创建页脚链接
func (website *WebsiteApi) WebsiteCreateFooterLink(c *gin.Context) {
	var req database.FooterLink
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	err = websiteService.WebsiteCreateFooterLink(req)

	if err != nil {
		global.Log.Error("Failed to create footer link:", zap.Error(err))
		response.FailWithMessage("Failed to create footer link", c)
		return
	}
	response.OkWithMessage("Successfully created footer link", c)
}

// WebsiteDeleteFooterLink 删除页脚链接
func (website *WebsiteApi) WebsiteDeleteFooterLink(c *gin.Context) {
	var req database.FooterLink
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	err = websiteService.WebsiteDeleteFooterLink(req)

	if err != nil {
		global.Log.Error("Failed to delete footer link:", zap.Error(err))
		response.FailWithMessage("Failed to delete footer link", c)
		return
	}
	response.OkWithMessage("Successfully deleted footer link", c)
}
