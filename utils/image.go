package utils

import (
	"server/model/appTypes"
	"server/model/database"

	"gorm.io/gorm"
)

// InitImagesCategory 初始化图片类别
func InitImagesCategory(tx *gorm.DB, urls []string) error {
	// 图片不是本地或七牛，不会报错。如果 urls 中有不匹配的记录（即数据库中没有对应的 URL），gorm 仍然会正常执行查询，不会因为某个 URL 没有匹配的记录而报错。
	return tx.Model(&database.Image{}).Where("url IN ?", urls).Update("category", appTypes.Null).Error
}

// ChangeImagesCategory 修改图片类别
func ChangeImagesCategory(tx *gorm.DB, urls []string, category appTypes.Category) error {
	return tx.Model(&database.Image{}).Where("url IN ?", urls).Update("category", category).Error
}
