package config

// Gaode 高德服务配置，详情请见 https://lbs.amap.com/
type Gaode struct {
	Enable bool   `json:"enable" yaml:"enable"` // 是否开启高德服务，true 表示启用，false 表示禁用
	Key    string `json:"key" yaml:"key"`       // 高德服务的应用密钥，用于身份验证和服务访问
}
