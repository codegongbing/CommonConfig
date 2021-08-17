package upload

import (
	"CommonConfig/global"
	"mime/multipart"
)

//@author: [ccfish86](https://github.com/ccfish86)
//@author: [SliverHorn](https://github.com/SliverHorn)
//@interface_name: OSS
//@description: OSS接口

type OSS interface {
	UploadFile(fileHeader *multipart.FileHeader, path ...string) (string, error)
	DeleteFile(key string) error
}

//@author: [ccfish86](https://github.com/ccfish86)
//@author: [SliverHorn](https://github.com/SliverHorn)
//@function: NewOss
//@description: OSS接口
//@description: OSS的实例化方法
//@return: OSS

func NewOSS() OSS {
	switch global.CONFIG.Project.OSSType {
	case "local":
		return &Local{}
	case "qiniu":
		return &QiNiu{}
	//case "tencent-cos":
	//	return &TencentCOS{}
	case "aliyun-oss":
		return &AliyunOSS{}
	default:
		return &QiNiu{}
	}
}
