package global

import (
	"time"

	"gorm.io/gorm"
)

type MODEL struct {
	ID        uint           `json:"id" gorm:"primarykey"` // 主键 ID
	CreatedAt time.Time      `json:"created_at"`           // 创建时间
	UpdatedAt time.Time      `json:"updated_at"`           // 更新时间
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`       // 删除时间
}
