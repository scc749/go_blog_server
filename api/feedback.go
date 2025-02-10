package api

import (
	"server/global"
	"server/model/request"
	"server/model/response"
	"server/utils"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type FeedbackApi struct {
}

// FeedbackNew 获取最新反馈
func (feedbackApi *FeedbackApi) FeedbackNew(c *gin.Context) {
	list, err := feedbackService.FeedbackNew()
	if err != nil {
		global.Log.Error("Failed to get new feedback:", zap.Error(err))
		response.FailWithMessage("Failed to get new feedback", c)
		return
	}
	response.OkWithData(list, c)
}

// FeedbackCreate 创建反馈
func (feedbackApi *FeedbackApi) FeedbackCreate(c *gin.Context) {
	var req request.FeedbackCreate
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	req.UUID = utils.GetUUID(c)
	err = feedbackService.FeedbackCreate(req)
	if err != nil {
		global.Log.Error("Failed to create feedback:", zap.Error(err))
		response.FailWithMessage("Failed to create feedback", c)
		return
	}
	response.OkWithMessage("Successfully created feedback", c)
}

// FeedbackInfo 获取用户反馈信息
func (feedbackApi *FeedbackApi) FeedbackInfo(c *gin.Context) {
	uuid := utils.GetUUID(c)
	list, err := feedbackService.FeedbackInfo(uuid)
	if err != nil {
		global.Log.Error("Failed to get feedback information:", zap.Error(err))
		response.FailWithMessage("Failed to get feedback information", c)
		return
	}
	response.OkWithData(list, c)
}

// FeedbackDelete 删除反馈
func (feedbackApi *FeedbackApi) FeedbackDelete(c *gin.Context) {
	var req request.FeedbackDelete
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	err = feedbackService.FeedbackDelete(req)
	if err != nil {
		global.Log.Error("Failed to delete feedback:", zap.Error(err))
		response.FailWithMessage("Failed to delete feedback", c)
		return
	}
	response.OkWithMessage("Successfully deleted feedback", c)
}

// FeedbackReply 回复反馈
func (feedbackApi *FeedbackApi) FeedbackReply(c *gin.Context) {
	var req request.FeedbackReply
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	err = feedbackService.FeedbackReply(req)
	if err != nil {
		global.Log.Error("Failed to update feedback:", zap.Error(err))
		response.FailWithMessage("Failed to update feedback", c)
		return
	}
	response.OkWithMessage("Successfully updated feedback", c)
}

// FeedbackList 获取反馈列表
func (feedbackApi *FeedbackApi) FeedbackList(c *gin.Context) {
	var pageInfo request.PageInfo
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	list, total, err := feedbackService.FeedbackList(pageInfo)
	if err != nil {
		global.Log.Error("Failed to get feedback list:", zap.Error(err))
		response.FailWithMessage("Failed to get feedback list", c)
		return
	}
	response.OkWithData(response.PageResult{
		List:  list,
		Total: total,
	}, c)
}
