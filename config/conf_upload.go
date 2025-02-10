package config

type Upload struct {
	Size int    `json:"size" yaml:"size"` // 图片上传的大小，单位 MB
	Path string `json:"path" yaml:"path"` // 图片上传的目录
}
