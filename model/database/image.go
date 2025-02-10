package database

import (
	"server/global"
	"server/model/appTypes"
)

// Image 图片表
type Image struct {
	global.MODEL
	Name     string            `json:"name"`                       // 名称
	URL      string            `json:"url" gorm:"size:255;unique"` // 路径
	Category appTypes.Category `json:"category"`                   // 类别
	Storage  appTypes.Storage  `json:"storage"`                    // 存储类型
}
