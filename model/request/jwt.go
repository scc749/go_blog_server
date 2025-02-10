package request

import (
	"server/model/appTypes"

	"github.com/gofrs/uuid"
	jwt "github.com/golang-jwt/jwt/v4"
)

// JwtCustomClaims 结构体用于存储JWT的自定义Claims，继承自BaseClaims，并包含标准的JWT注册信息
type JwtCustomClaims struct {
	BaseClaims           // 基础Claims，包含用户ID、UUID和角色ID
	jwt.RegisteredClaims // 标准JWT声明，例如过期时间、发行者等
}

// JwtCustomRefreshClaims 结构体用于存储刷新Token的自定义Claims，包含用户ID和标准的JWT注册信息
type JwtCustomRefreshClaims struct {
	UserID               uint // 用户ID，用于与刷新Token相关的身份验证
	jwt.RegisteredClaims      // 标准JWT声明
}

// BaseClaims 结构体用于存储基本的用户信息，作为JWT的Claim部分
type BaseClaims struct {
	UserID uint            // 用户ID，标识用户唯一性
	UUID   uuid.UUID       // 用户的UUID，唯一标识用户
	RoleID appTypes.RoleID // 用户角色ID，表示用户的权限级别
}
