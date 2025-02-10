package utils

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/url"
)

// HttpRequest 函数用于发送HTTP请求
func HttpRequest(
	urlStr string, // 请求的URL字符串
	method string, // 请求方法（GET, POST等）
	headers map[string]string, // 请求头（如 Content-Type 等）
	params map[string]string, // 查询参数（如 ?key=value&key2=value2）
	data any) (*http.Response, error) { // 请求体的内容（如果有的话）

	// 创建URL对象
	u, err := url.Parse(urlStr) // 将urlStr解析为URL对象
	if err != nil {
		return nil, err // 如果URL解析失败，返回错误
	}

	// 向URL添加查询参数
	query := u.Query()         // 获取URL中的查询部分（如果有）
	for k, v := range params { // 遍历参数并添加到URL中
		query.Set(k, v) // 使用Set方法保证参数键值对唯一
	}
	u.RawQuery = query.Encode() // 更新URL的查询部分

	// 将请求体数据（如果有）编码成JSON格式
	buf := new(bytes.Buffer) // 创建一个缓冲区用于存储请求体
	if data != nil {
		b, err := json.Marshal(data) // 将data编码为JSON字节数组
		if err != nil {
			return nil, err // 如果编码失败，返回错误
		}
		buf = bytes.NewBuffer(b) // 将编码后的字节数组转换为缓冲区
	}

	// 创建HTTP请求对象
	req, err := http.NewRequest(method, u.String(), buf) // 使用指定的URL和方法创建请求
	if err != nil {
		return nil, err // 如果请求创建失败，返回错误
	}

	// 设置请求头
	for k, v := range headers { // 遍历传入的头部信息并设置到请求中
		req.Header.Set(k, v) // 设置头部
	}

	// 如果请求体存在，将Content-Type设置为application/json
	if data != nil {
		req.Header.Set("Content-Type", "application/json") // 设置请求头为JSON类型
	}

	// 发送HTTP请求并获取响应
	resp, err := http.DefaultClient.Do(req) // 使用默认的HTTP客户端发送请求
	if err != nil {
		return nil, err // 如果请求失败，返回错误
	}

	// 返回响应对象
	return resp, nil // 返回响应，调用者可根据需要处理响应数据
}
