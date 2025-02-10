package database

// ArticleCategory 文章类别表
type ArticleCategory struct {
	Category string `json:"category" gorm:"primaryKey"` // 类别
	Number   int    `json:"number"`                     // 统计数量
}
