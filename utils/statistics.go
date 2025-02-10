package utils

import "gorm.io/gorm"

// FetchDateCounts 用于根据查询条件获取日期统计数据
func FetchDateCounts(db *gorm.DB, query *gorm.DB) map[string]int {
	var dateCounts []struct {
		Date  string `json:"date"`
		Count int    `json:"count"`
	}
	db.Where(query).
		Select("date_format(created_at, '%Y-%m-%d') as date", "count(id) as count").
		Group("date").Scan(&dateCounts)

	dateCountMap := make(map[string]int)
	for _, count := range dateCounts {
		dateCountMap[count.Date] = count.Count
	}
	return dateCountMap
}
