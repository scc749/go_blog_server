package task

import (
	"server/global"

	"github.com/robfig/cron/v3"
	"go.uber.org/zap"
)

func RegisterScheduledTasks(c *cron.Cron) error {
	if _, err := c.AddFunc("@hourly", func() {
		if err := UpdateArticleViewsSyncTask(); err != nil {
			global.Log.Error("Failed to update article views:", zap.Error(err))
		}
	}); err != nil {
		return err
	}
	if _, err := c.AddFunc("@hourly", func() {
		if err := GetHotListSyncTask(); err != nil {
			global.Log.Error("Failed to get hot list:", zap.Error(err))
		}
	}); err != nil {
		return err
	}
	if _, err := c.AddFunc("@daily", func() {
		if err := GetCalendarSyncTask(); err != nil {
			global.Log.Error("Failed to get calendar:", zap.Error(err))
		}
	}); err != nil {
		return err
	}
	return nil
}
