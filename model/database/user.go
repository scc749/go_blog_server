package database

import (
	"server/global"
	"server/model/appTypes"

	"github.com/gofrs/uuid"
)

// User 用户表
type User struct {
	global.MODEL
	UUID      uuid.UUID         `json:"uuid" gorm:"type:char(36);unique"`              // uuid
	Username  string            `json:"username"`                                      // 用户名
	Password  string            `json:"-"`                                             // 密码
	Email     string            `json:"email"`                                         // 邮箱
	Openid    string            `json:"openid"`                                        // openid
	Avatar    string            `json:"avatar" gorm:"size:255"`                        // 头像：邮箱注册的头像或 QQ 登录的空间头像
	Address   string            `json:"address"`                                       // 地址
	Signature string            `json:"signature" gorm:"default:'签名是空白的，这位用户似乎比较低调。'"` // 签名
	RoleID    appTypes.RoleID   `json:"role_id"`                                       // 角色 ID
	Register  appTypes.Register `json:"register"`                                      // 注册来源
	Freeze    bool              `json:"freeze"`                                        // 用户是否被冻结
}
