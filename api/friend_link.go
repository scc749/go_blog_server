package api

import (
	"server/global"
	"server/model/request"
	"server/model/response"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type FriendLinkApi struct {
}

// FriendLinkInfo 获取友链信息
func (friendLinkApi *FriendLinkApi) FriendLinkInfo(c *gin.Context) {
	list, total, err := friendLinkService.FriendLinkInfo()
	if err != nil {
		global.Log.Error("Failed to get friend link information:", zap.Error(err))
		response.FailWithMessage("Failed to get friend link information", c)
		return
	}
	response.OkWithData(response.FriendLinkInfo{
		List:  list,
		Total: total,
	}, c)
}

// FriendLinkCreate 创建友链
func (friendLinkApi *FriendLinkApi) FriendLinkCreate(c *gin.Context) {
	var req request.FriendLinkCreate
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	err = friendLinkService.FriendLinkCreate(req)
	if err != nil {
		global.Log.Error("Failed to create friend link:", zap.Error(err))
		response.FailWithMessage("Failed to create friend link", c)
		return
	}
	response.OkWithMessage("Successfully created friend link", c)
}

// FriendLinkDelete 删除友链
func (friendLinkApi *FriendLinkApi) FriendLinkDelete(c *gin.Context) {
	var req request.FriendLinkDelete
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	err = friendLinkService.FriendLinkDelete(req)
	if err != nil {
		global.Log.Error("Failed to delete friend link:", zap.Error(err))
		response.FailWithMessage("Failed to delete friend link", c)
		return
	}
	response.OkWithMessage("Successfully deleted friend link", c)
}

// FriendLinkUpdate 更新友链
func (friendLinkApi *FriendLinkApi) FriendLinkUpdate(c *gin.Context) {
	var req request.FriendLinkUpdate
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	err = friendLinkService.FriendLinkUpdate(req)
	if err != nil {
		global.Log.Error("Failed to update friend link:", zap.Error(err))
		response.FailWithMessage("Failed to update friend link", c)
		return
	}
	response.OkWithMessage("Successfully updated friend link", c)
}

// FriendLinkList 获取友链列表
func (friendLinkApi *FriendLinkApi) FriendLinkList(c *gin.Context) {
	var pageInfo request.FriendLinkList
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	list, total, err := friendLinkService.FriendLinkList(pageInfo)
	if err != nil {
		global.Log.Error("Failed to get friend link list:", zap.Error(err))
		response.FailWithMessage("Failed to get friend link list", c)
		return
	}
	response.OkWithData(response.PageResult{
		List:  list,
		Total: total,
	}, c)
}
