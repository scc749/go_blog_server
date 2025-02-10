package database

import (
	"server/global"

	"github.com/gofrs/uuid"
)

// Feedback 反馈表
type Feedback struct {
	global.MODEL
	UserUUID uuid.UUID `json:"user_uuid" gorm:"type:char(36)"`               // 用户 uuid
	User     User      `json:"-" gorm:"foreignKey:UserUUID;references:UUID"` // 关联的用户
	Content  string    `json:"content"`                                      // 内容
	Reply    string    `json:"reply"`                                        // 回复
}
