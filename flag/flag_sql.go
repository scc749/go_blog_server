package flag

import (
	"server/global"
	"server/model/database"
)

// SQL 表结构迁移，如果表不存在，它会创建新表；如果表已经存在，它会根据结构更新表
func SQL() error {
	return global.DB.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(
		&database.Advertisement{},
		&database.ArticleCategory{},
		&database.ArticleLike{},
		&database.ArticleTag{},
		&database.Comment{},
		&database.Feedback{},
		&database.FooterLink{},
		&database.FriendLink{},
		&database.Image{},
		&database.JwtBlacklist{},
		&database.Login{},
		&database.User{},
	)
}
