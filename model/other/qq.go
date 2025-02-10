package other

// AccessTokenResponse 表示通过授权码获取的 Access Token 返回结构
type AccessTokenResponse struct {
	AccessToken  string `json:"access_token"`  // 授权令牌
	ExpiresIn    string `json:"expires_in"`    // 该 access token 的有效期，单位为秒
	RefreshToken string `json:"refresh_token"` // 刷新 token
	Openid       string `json:"openid"`        // 用户的 Openid
}

// UserInfoResponse 表示获取用户信息的返回结构
type UserInfoResponse struct {
	Ret          int    `json:"ret"`            // 返回码
	Msg          string `json:"msg"`            // 错误信息
	IsLost       int    `json:"is_lost"`        // 数据丢失标识
	Nickname     string `json:"nickname"`       // 用户昵称
	Figureurl    string `json:"figureurl"`      // 30x30头像URL
	Figureurl1   string `json:"figureurl_1"`    // 50x50头像URL
	Figureurl2   string `json:"figureurl_2"`    // 100x100头像URL
	FigureurlQQ1 string `json:"figureurl_qq_1"` // 40x40 QQ头像URL
	FigureurlQQ2 string `json:"figureurl_qq_2"` // 100x100 QQ头像URL
}
