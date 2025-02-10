package utils

import (
	"net"
	"server/global"
	"server/model/appTypes"
	"server/model/request"

	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
	"go.uber.org/zap"
)

// SetRefreshToken 设置Refresh Token的cookie
func SetRefreshToken(c *gin.Context, token string, maxAge int) {
	// 获取请求的host，如果失败则取原始请求host
	host, _, err := net.SplitHostPort(c.Request.Host)
	if err != nil {
		host = c.Request.Host
	}
	// 调用setCookie设置refresh-token
	setCookie(c, "x-refresh-token", token, maxAge, host)
}

// ClearRefreshToken 清除Refresh Token的cookie
func ClearRefreshToken(c *gin.Context) {
	// 获取请求的host，如果失败则取原始请求host
	host, _, err := net.SplitHostPort(c.Request.Host)
	if err != nil {
		host = c.Request.Host
	}
	// 调用setCookie设置cookie值为空并过期，删除refresh-token
	setCookie(c, "x-refresh-token", "", -1, host)
}

// setCookie 设置指定名称和值的cookie
func setCookie(c *gin.Context, name, value string, maxAge int, host string) {
	// 判断host是否是IP地址
	if net.ParseIP(host) != nil {
		// 如果是IP地址，设置cookie的domain为“/”
		c.SetCookie(name, value, maxAge, "/", "", false, true)
	} else {
		// 如果是域名，设置cookie的domain为域名
		c.SetCookie(name, value, maxAge, "/", host, false, true)
	}
}

// GetAccessToken 从请求头获取Access Token
func GetAccessToken(c *gin.Context) string {
	// 获取x-access-token头部值
	token := c.Request.Header.Get("x-access-token")
	return token
}

// GetRefreshToken 从cookie获取Refresh Token
func GetRefreshToken(c *gin.Context) string {
	// 尝试从cookie中获取refresh-token
	token, _ := c.Cookie("x-refresh-token")
	return token
}

// GetClaims 从Gin的Context中解析并获取JWT的Claims
func GetClaims(c *gin.Context) (*request.JwtCustomClaims, error) {
	// 获取请求头中的Access Token
	token := GetAccessToken(c)
	// 创建JWT实例
	j := NewJWT()
	// 解析Access Token
	claims, err := j.ParseAccessToken(token)
	if err != nil {
		// 如果解析失败，记录错误日志
		global.Log.Error("Failed to retrieve JWT parsing information from Gin's Context. Please check if the request header contains 'x-access-token' and if the claims structure is correct.", zap.Error(err))
	}
	return claims, err
}

// GetRefreshClaims 从Gin的Context中解析并获取Refresh Token的Claims
func GetRefreshClaims(c *gin.Context) (*request.JwtCustomRefreshClaims, error) {
	// 获取Refresh Token
	token := GetRefreshToken(c)
	// 创建JWT实例
	j := NewJWT()
	// 解析Refresh Token
	claims, err := j.ParseRefreshToken(token)
	if err != nil {
		// 如果解析失败，记录错误日志
		global.Log.Error("Failed to retrieve JWT parsing information from Gin's Context. Please check if the request header contains 'x-refresh-token' and if the claims structure is correct.", zap.Error(err))
	}
	return claims, err
}

// GetUserInfo 从Gin的Context中获取JWT解析出来的用户信息（Claims）
func GetUserInfo(c *gin.Context) *request.JwtCustomClaims {
	// 首先尝试从Context中获取"claims"
	if claims, exists := c.Get("claims"); !exists {
		// 如果不存在，则重新解析Access Token
		if cl, err := GetClaims(c); err != nil {
			// 如果解析失败，返回nil
			return nil
		} else {
			// 返回解析出来的用户信息
			return cl
		}
	} else {
		// 如果已存在claims，则直接返回
		waitUse := claims.(*request.JwtCustomClaims)
		return waitUse
	}
}

// GetUserID 从Gin的Context中获取JWT解析出来的用户ID
func GetUserID(c *gin.Context) uint {
	// 首先尝试从Context中获取"claims"
	if claims, exists := c.Get("claims"); !exists {
		// 如果不存在，则重新解析Access Token
		if cl, err := GetClaims(c); err != nil {
			// 如果解析失败，返回0
			return 0
		} else {
			// 返回解析出来的用户ID
			return cl.UserID
		}
	} else {
		// 如果已存在claims，则直接返回用户ID
		waitUse := claims.(*request.JwtCustomClaims)
		return waitUse.UserID
	}
}

// GetUUID 从Gin的Context中获取JWT解析出来的用户UUID
func GetUUID(c *gin.Context) uuid.UUID {
	// 首先尝试从Context中获取"claims"
	if claims, exists := c.Get("claims"); !exists {
		// 如果不存在，则重新解析Access Token
		if cl, err := GetClaims(c); err != nil {
			// 如果解析失败，返回一个空UUID
			return uuid.UUID{}
		} else {
			// 返回解析出来的UUID
			return cl.UUID
		}
	} else {
		// 如果已存在claims，则直接返回UUID
		waitUse := claims.(*request.JwtCustomClaims)
		return waitUse.UUID
	}
}

// GetRoleID 从Gin的Context中获取JWT解析出来的用户角色ID
func GetRoleID(c *gin.Context) appTypes.RoleID {
	// 首先尝试从Context中获取"claims"
	if claims, exists := c.Get("claims"); !exists {
		// 如果不存在，则重新解析Access Token
		if cl, err := GetClaims(c); err != nil {
			// 如果解析失败，返回0
			return 0
		} else {
			// 返回解析出来的角色ID
			return cl.RoleID
		}
	} else {
		// 如果已存在claims，则直接返回角色ID
		waitUse := claims.(*request.JwtCustomClaims)
		return waitUse.RoleID
	}
}
