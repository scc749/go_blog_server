package service

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"server/global"
	"server/model/other"
	"server/utils"
)

type QQService struct {
}

var QQServiceAPP = new(QQService)

// GetAccessTokenByCode 通过Authorization Code获取Access Token
func (qqService *QQService) GetAccessTokenByCode(code string) (other.AccessTokenResponse, error) {
	data := other.AccessTokenResponse{}
	clientID := global.Config.QQ.AppID
	clientSecret := global.Config.QQ.AppKey
	redirectUri := global.Config.QQ.RedirectURI
	urlStr := "https://graph.qq.com/oauth2.0/token"
	method := "GET"
	params := map[string]string{
		"grant_type":    "authorization_code",
		"client_id":     clientID,
		"client_secret": clientSecret,
		"code":          code,
		"redirect_uri":  redirectUri,
		"fmt":           "json",
		"need_openid":   "1",
	}
	res, err := utils.HttpRequest(urlStr, method, nil, params, nil)
	if err != nil {
		return data, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return data, fmt.Errorf("request failed with status code: %d", res.StatusCode)
	}

	byteData, err := io.ReadAll(res.Body)
	if err != nil {
		return data, err
	}

	err = json.Unmarshal(byteData, &data)
	if err != nil {
		return data, err
	}
	return data, nil
}

// GetUserInfoByAccessTokenAndOpenid 获取登录用户信息
func (qqService *QQService) GetUserInfoByAccessTokenAndOpenid(accessToken, openID string) (other.UserInfoResponse, error) {
	data := other.UserInfoResponse{}
	oauthConsumerKey := global.Config.QQ.AppID
	urlStr := "https://graph.qq.com/user/get_user_info"
	method := "GET"
	params := map[string]string{
		"access_token":       accessToken,
		"oauth_consumer_key": oauthConsumerKey,
		"openid":             openID,
	}
	res, err := utils.HttpRequest(urlStr, method, nil, params, nil)
	if err != nil {
		return data, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return data, fmt.Errorf("request failed with status code: %d", res.StatusCode)
	}

	byteData, err := io.ReadAll(res.Body)
	if err != nil {
		return data, err
	}

	err = json.Unmarshal(byteData, &data)
	if err != nil {
		return data, err
	}
	return data, nil
}
