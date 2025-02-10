package config

// Qiniu 七牛云配置，详情请见 https://www.qiniu.com/
type Qiniu struct {
	Zone          string `json:"zone" yaml:"zone"`                       // 存储区域
	Bucket        string `json:"bucket" yaml:"bucket"`                   // 空间名称
	ImgPath       string `json:"img_path" yaml:"img_path"`               // CDN 加速域名
	AccessKey     string `json:"access_key" yaml:"access_key"`           // 秘钥 AK
	SecretKey     string `json:"secret_key" yaml:"secret_key"`           // 秘钥 SK
	UseHTTPS      bool   `json:"use_https" yaml:"use_https"`             // 是否使用 https
	UseCdnDomains bool   `json:"use_cdn_domains" yaml:"use_cdn_domains"` // 上传是否使用 CDN 上传加速
}
