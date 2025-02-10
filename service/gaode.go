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

// JwtService 提供与高德相关的服务
type GaodeService struct {
}

// GetLocationByIP 根据IP地址获取地理位置信息
func (gaodeService *GaodeService) GetLocationByIP(ip string) (other.IPResponse, error) {
	data := other.IPResponse{}
	key := global.Config.Gaode.Key
	urlStr := "https://restapi.amap.com/v3/ip"
	method := "GET"
	params := map[string]string{
		"ip":  ip,
		"key": key,
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

// GetWeatherByAdcode 根据城市编码获取实时天气信息
func (gaodeService *GaodeService) GetWeatherByAdcode(adcode string) (other.Live, error) {
	data := other.WeatherResponse{}
	key := global.Config.Gaode.Key
	urlStr := "https://restapi.amap.com/v3/weather/weatherInfo"
	method := "GET"
	params := map[string]string{
		"city": adcode,
		"key":  key,
	}
	res, err := utils.HttpRequest(urlStr, method, nil, params, nil)
	if err != nil {
		return other.Live{}, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return other.Live{}, fmt.Errorf("request failed with status code: %d", res.StatusCode)
	}

	byteData, err := io.ReadAll(res.Body)
	if err != nil {
		return other.Live{}, err
	}

	err = json.Unmarshal(byteData, &data)
	if err != nil {
		return other.Live{}, err
	}

	// 检查是否有返回的天气数据
	if len(data.Lives) == 0 {
		return other.Live{}, fmt.Errorf("no live weather data available") // 没有天气数据时返回错误
	}

	// 返回当天的天气数据
	return data.Lives[0], nil
}
