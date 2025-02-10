package service

import (
	"encoding/json"
	"server/global"
	"server/model/other"
	"server/utils/hotSearch"
	"time"
)

type HotSearchService struct {
}

func (hotSearchService *HotSearchService) GetHotSearchDataBySource(sourceStr string) (other.HotSearchData, error) {
	result, err := global.Redis.Get(sourceStr).Result()
	if err != nil {
		source := hotSearch.NewSource(sourceStr)
		hotSearchData, err := source.GetHotSearchData(30)
		if err != nil {
			return other.HotSearchData{}, err
		}
		bytes, err := json.Marshal(hotSearchData)
		if err != nil {
			return other.HotSearchData{}, err
		}
		if err := global.Redis.Set(sourceStr, bytes, time.Hour).Err(); err != nil {
			return other.HotSearchData{}, err
		}
		return hotSearchData, nil
	}
	var hotSearchData other.HotSearchData
	if err := json.Unmarshal([]byte(result), &hotSearchData); err != nil {
		return other.HotSearchData{}, err
	}
	return hotSearchData, nil
}
