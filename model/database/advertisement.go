package database

import "server/global"

// Advertisement 广告表
type Advertisement struct {
	global.MODEL
	AdImage string `json:"ad_image" gorm:"size:255"` // 图片
	Image   Image  `json:"-" gorm:"foreignKey:AdImage;references:URL"`
	Link    string `json:"link"`    // 链接
	Title   string `json:"title"`   // 标题
	Content string `json:"content"` // 内容
}
