package initialize

import (
	"os"
	"server/global"
	"server/utils"

	"github.com/songzhibin97/gkit/cache/local_cache"
	"go.uber.org/zap"
)

// OtherInit 执行其他配置初始化
func OtherInit() {
	// 解析刷新令牌过期时间
	refreshTokenExpiry, err := utils.ParseDuration(global.Config.Jwt.RefreshTokenExpiryTime)
	if err != nil {
		global.Log.Error("Failed to parse refresh token expiry time configuration:", zap.Error(err))
		os.Exit(1)
	}

	// 解析访问令牌过期时间
	_, err = utils.ParseDuration(global.Config.Jwt.AccessTokenExpiryTime)
	if err != nil {
		global.Log.Error("Failed to parse access token expiry time configuration:", zap.Error(err))
		os.Exit(1)
	}

	// 配置本地缓存过期时间（使用刷新令牌过期时间，方便在远程登录或账户冻结时对 JWT 进行黑名单处理）
	global.BlackCache = local_cache.NewCache(
		local_cache.SetDefaultExpire(refreshTokenExpiry),
	)
}
