package upload

import (
	"CommonConfig/global"
	"CommonConfig/utils"
	"errors"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"go.uber.org/zap"
	"mime/multipart"
	"strings"
)

type AliyunOSS struct{}

func (*AliyunOSS) UploadFile(fileHeader *multipart.FileHeader, path ...string) (string, error) {
	file, err := fileHeader.Open()
	if err != nil {
		global.LOG.Error("function fileHeader.Open() Filed", zap.Any("err", err.Error()))
		return "", errors.New("function fileHeader.Open() Filed, err:" + err.Error())
	}
	defer func(file multipart.File) {
		err := file.Close()
		if err != nil {
			global.LOG.Error("function file.Close() Filed", zap.Any("err", err.Error()))
		}
	}(file)
	bucket, err := NewBucket()
	if err != nil {
		global.LOG.Error("New AliyunOSS Bucket Failed", zap.Any("err", err.Error()))
		return "", errors.New("New AliyunOSS Bucket Failed, err:" + err.Error())
	}
	//生成随机文件名
	fileName := utils.RandomFileName(fileHeader.Filename, path...)
	err = bucket.PutObject(fileName, file)
	if err != nil {
		global.LOG.Error("AliyunOSS upload file Failed", zap.Any("err", err.Error()))
		return "", errors.New("AliyunOSS upload file Failed, err:" + err.Error())
	}

	return global.CONFIG.AliyunOSS.URL + "/" + fileName, nil
}

func (*AliyunOSS) DeleteFile(key string) error {
	key = strings.Replace(key, global.CONFIG.AliyunOSS.URL+"/", "", 1)
	bucket, err := NewBucket()
	if err != nil {
		global.LOG.Error("New AliyunOSS Bucket Failed", zap.Any("err", err.Error()))
		return errors.New("New AliyunOSS Bucket Failed, err:" + err.Error())
	}
	// 删除单个文件。objectName表示删除OSS文件时需要指定包含文件后缀在内的完整路径，例如abc/efg/123.jpg。
	// 如需删除文件夹，请将objectName设置为对应的文件夹名称。如果文件夹非空，则需要将文件夹下的所有object删除后才能删除该文件夹。
	err = bucket.DeleteObject(key)
	if err != nil {
		global.LOG.Error("AliyunOSS delete file Filed", zap.Any("err", err.Error()))
		return errors.New("AliyunOSS delete file Filed, err:" + err.Error())
	}
	return nil
}

func NewBucket() (*oss.Bucket, error) {
	// 创建OSSClient实例。
	client, err := oss.New(
		global.CONFIG.AliyunOSS.Endpoint, global.CONFIG.AliyunOSS.AccessKeyId,
		global.CONFIG.AliyunOSS.AccessKeySecret,
	)
	if err != nil {
		return nil, err
	}

	// 获取存储空间。
	bucket, err := client.Bucket(global.CONFIG.AliyunOSS.BucketName)
	if err != nil {
		return nil, err
	}

	return bucket, nil
}
