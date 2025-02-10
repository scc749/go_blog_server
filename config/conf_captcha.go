package config

type Captcha struct {
	Height   int     `json:"height" yaml:"height"`       // PNG 图片的高度（以像素为单位）
	Width    int     `json:"width" yaml:"width"`         // 验证码 PNG 图片的宽度（以像素为单位）
	Length   int     `json:"length" yaml:"length"`       // 验证码结果中默认的数字个数
	MaxSkew  float64 `json:"max_skew" yaml:"max_skew"`   // 单个数字的最大偏斜因子（绝对值）
	DotCount int     `json:"dot_count" yaml:"dot_count"` // 背景圆点的数量
}
