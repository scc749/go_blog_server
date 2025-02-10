package database

// FooterLink 页脚链接表
type FooterLink struct {
	Title string `json:"title" gorm:"primaryKey"` // 标题
	Link  string `json:"link"`                    // 链接
}
