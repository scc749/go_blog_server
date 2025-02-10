package database

// ArticleTag 文章标签表
type ArticleTag struct {
	Tag    string `json:"tag" gorm:"primaryKey"` // 标签
	Number int    `json:"number"`                // 统计数量
}
