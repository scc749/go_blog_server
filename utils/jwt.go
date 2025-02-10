package utils

import (
	"errors"
	"server/global"
	"server/model/request"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type JWT struct {
	AccessTokenSecret  []byte
	RefreshTokenSecret []byte
}

var (
	TokenExpired     = errors.New("token is expired")
	TokenNotValidYet = errors.New("token not active yet")
	TokenMalformed   = errors.New("that's not even a token")
	TokenInvalid     = errors.New("couldn't handle this token")
)

func NewJWT() *JWT {
	return &JWT{
		AccessTokenSecret:  []byte(global.Config.Jwt.AccessTokenSecret),
		RefreshTokenSecret: []byte(global.Config.Jwt.RefreshTokenSecret),
	}
}

// CreateAccessClaims 创建 Access Token 的 Claims，包含基本信息和过期时间等
func (j *JWT) CreateAccessClaims(baseClaims request.BaseClaims) request.JwtCustomClaims {
	ep, _ := ParseDuration(global.Config.Jwt.AccessTokenExpiryTime) // 获取过期时间
	claims := request.JwtCustomClaims{
		BaseClaims: baseClaims, // 基本 Claims
		RegisteredClaims: jwt.RegisteredClaims{
			Audience:  jwt.ClaimStrings{"TAP"},                // 受众
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(ep)), // 过期时间
			Issuer:    global.Config.Jwt.Issuer,               // 签名的发行者
		},
	}
	return claims
}

// CreateAccessToken 创建 Access Token，通过 Claims 生成 JWT Token
func (j *JWT) CreateAccessToken(claims request.JwtCustomClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims) // 创建新的 JWT Token
	return token.SignedString(j.AccessTokenSecret)             // 使用 AccessToken 密钥签名并返回 Token 字符串
}

// CreateRefreshClaims 创建 Refresh Token 的 Claims，包含用户信息和过期时间等
func (j *JWT) CreateRefreshClaims(baseClaims request.BaseClaims) request.JwtCustomRefreshClaims {
	ep, _ := ParseDuration(global.Config.Jwt.RefreshTokenExpiryTime) // 获取过期时间
	claims := request.JwtCustomRefreshClaims{
		UserID: baseClaims.UserID, // 用户 ID
		RegisteredClaims: jwt.RegisteredClaims{
			Audience:  jwt.ClaimStrings{"TAP"},                // 受众
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(ep)), // 过期时间
			Issuer:    global.Config.Jwt.Issuer,               // 签名的发行者
		},
	}
	return claims
}

// CreateRefreshToken 创建 Refresh Token，通过 Claims 生成 JWT Token
func (j *JWT) CreateRefreshToken(claims request.JwtCustomRefreshClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims) // 创建新的 JWT Token
	return token.SignedString(j.RefreshTokenSecret)            // 使用 RefreshToken 密钥签名并返回 Token 字符串
}

// ParseAccessToken 解析 Access Token，验证 Token 并返回 Claims 信息
func (j *JWT) ParseAccessToken(tokenString string) (*request.JwtCustomClaims, error) {
	claims, err := j.parseToken(tokenString, &request.JwtCustomClaims{}, j.AccessTokenSecret) // 解析 Token
	if err != nil {
		return nil, err
	}
	if customClaims, ok := claims.(*request.JwtCustomClaims); ok { // 确保解析出的 Claims 类型正确
		return customClaims, nil
	}
	return nil, TokenInvalid // 如果解析结果无效，返回 TokenInvalid 错误
}

// ParseRefreshToken 解析 Refresh Token，验证 Token 并返回 Claims 信息
func (j *JWT) ParseRefreshToken(tokenString string) (*request.JwtCustomRefreshClaims, error) {
	claims, err := j.parseToken(tokenString, &request.JwtCustomRefreshClaims{}, j.RefreshTokenSecret) // 解析 Token
	if err != nil {
		return nil, err
	}
	if refreshClaims, ok := claims.(*request.JwtCustomRefreshClaims); ok { // 确保解析出的 Claims 类型正确
		return refreshClaims, nil
	}
	return nil, TokenInvalid // 如果解析结果无效，返回 TokenInvalid 错误
}

// parseToken 通用的 Token 解析方法，验证 Token 是否有效并返回 Claims
func (j *JWT) parseToken(tokenString string, claims jwt.Claims, secretKey interface{}) (interface{}, error) {
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil // 返回密钥以验证 Token
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok { // 处理 Token 验证错误
			switch {
			case ve.Errors&jwt.ValidationErrorMalformed != 0:
				return nil, TokenMalformed // Token 格式错误
			case ve.Errors&jwt.ValidationErrorExpired != 0:
				return nil, TokenExpired // Token 已过期
			case ve.Errors&jwt.ValidationErrorNotValidYet != 0:
				return nil, TokenNotValidYet // Token 还未生效
			default:
				return nil, TokenInvalid // 其他错误返回 Token 无效
			}
		}
		return nil, TokenInvalid // 默认返回 Token 无效错误
	}

	if token.Valid { // 如果 Token 验证通过，返回 Claims
		return token.Claims, nil
	}
	return nil, TokenInvalid // Token 无效，返回错误
}
