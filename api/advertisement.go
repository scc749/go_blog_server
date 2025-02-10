package api

import (
	"server/global"
	"server/model/request"
	"server/model/response"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type AdvertisementApi struct {
}

// AdvertisementInfo 获取广告信息
func (advertisementApi *AdvertisementApi) AdvertisementInfo(c *gin.Context) {
	list, total, err := advertisementService.AdvertisementInfo()
	if err != nil {
		global.Log.Error("Failed to get advertisement information:", zap.Error(err))
		response.FailWithMessage("Failed to get advertisement information", c)
		return
	}
	response.OkWithData(response.AdvertisementInfo{
		List:  list,
		Total: total,
	}, c)
}

// AdvertisementCreate 创建广告
func (advertisementApi *AdvertisementApi) AdvertisementCreate(c *gin.Context) {
	var req request.AdvertisementCreate
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	err = advertisementService.AdvertisementCreate(req)
	if err != nil {
		global.Log.Error("Failed to create advertisement:", zap.Error(err))
		response.FailWithMessage("Failed to create advertisement", c)
		return
	}
	response.OkWithMessage("Successfully created advertisement", c)
}

// AdvertisementDelete 删除广告
func (advertisementApi *AdvertisementApi) AdvertisementDelete(c *gin.Context) {
	var req request.AdvertisementDelete
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	err = advertisementService.AdvertisementDelete(req)
	if err != nil {
		global.Log.Error("Failed to delete advertisement:", zap.Error(err))
		response.FailWithMessage("Failed to delete advertisement", c)
		return
	}
	response.OkWithMessage("Successfully deleted advertisement", c)
}

// AdvertisementUpdate 更新广告
func (advertisementApi *AdvertisementApi) AdvertisementUpdate(c *gin.Context) {
	var req request.AdvertisementUpdate
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	err = advertisementService.AdvertisementUpdate(req)
	if err != nil {
		global.Log.Error("Failed to update advertisement:", zap.Error(err))
		response.FailWithMessage("Failed to update advertisement", c)
		return
	}
	response.OkWithMessage("Successfully updated advertisement", c)
}

// AdvertisementList 获取广告列表
func (advertisementApi *AdvertisementApi) AdvertisementList(c *gin.Context) {
	var pageInfo request.AdvertisementList
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	list, total, err := advertisementService.AdvertisementList(pageInfo)
	if err != nil {
		global.Log.Error("Failed to get advertisement list:", zap.Error(err))
		response.FailWithMessage("Failed to get advertisement list", c)
		return
	}
	response.OkWithData(response.PageResult{
		List:  list,
		Total: total,
	}, c)
}
