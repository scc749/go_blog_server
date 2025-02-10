package upload

import (
	"mime/multipart"
	"server/global"
	"server/model/appTypes"
)

// WhiteImageList 定义一个白名单映射，包含支持的图片文件类型
var WhiteImageList = map[string]struct{}{
	".jpg":  {},
	".png":  {},
	".jpeg": {},
	".ico":  {},
	".tiff": {},
	".gif":  {},
	".svg":  {},
	".webp": {},
}

// OSS 对象存储接口定义，规定了文件上传和删除方法
type OSS interface {
	UploadImage(file *multipart.FileHeader) (string, string, error)
	DeleteImage(key string) error
}

// NewOss 是 OSS 的实例化方法，根据配置中的 OssType 来选择使用的存储类型
func NewOss() OSS {
	switch global.Config.System.OssType {
	case "local":
		return &Local{}
	case "qiniu":
		return &Qiniu{}
	default:
		return &Local{}
	}
}

// NewOssWithStorage 是根据传入的存储类型返回相应的 OSS 实例
func NewOssWithStorage(storage appTypes.Storage) OSS {
	switch storage {
	case appTypes.Local:
		return &Local{}
	case appTypes.Qiniu:
		return &Qiniu{}
	default:
		return &Local{}
	}
}
