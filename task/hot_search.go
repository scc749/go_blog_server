package task

import (
	"encoding/json"
	"server/global"
	"server/utils/hotSearch"
	"time"
)

func GetHotListSyncTask() error {
	sourceStrs := []string{"baidu", "zhihu", "kuaishou", "toutiao"}
	for _, sourceStr := range sourceStrs {
		source := hotSearch.NewSource(sourceStr)
		hotSearchData, err := source.GetHotSearchData(30)
		if err != nil {
			return err
		}
		data, err := json.Marshal(hotSearchData)
		if err != nil {
			return err
		}
		if err := global.Redis.Set(sourceStr, data, time.Hour).Err(); err != nil {
			return err
		}
	}
	return nil
}
