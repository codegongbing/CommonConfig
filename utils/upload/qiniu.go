package upload

import (
	"CommonConfig/global"
	"CommonConfig/utils"
	"context"
	"errors"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
	"go.uber.org/zap"
	"mime/multipart"
	"strings"
)

type QiNiu struct{}

//@author: [codegongbing]
//@object: *Qiniu
//@function: UploadFile
//@description: 上传文件
//@param: file *multipart.FileHeader, path ...string
//@return: string, error

func (*QiNiu) UploadFile(fileHeader *multipart.FileHeader, path ...string) (string, error) {
	file, err := fileHeader.Open()
	if err != nil {
		global.LOG.Error("function fileHeader.Open() Failed", zap.Any("err", err.Error()))
		return "", errors.New("function fileHeader.Open() Failed, err:" + err.Error())
	}
	defer func(file multipart.File) {
		err := file.Close()
		if err != nil {
			global.LOG.Error("function file.Close() Failed", zap.Any("err", err.Error()))
		}
	}(file)
	putPolicy := storage.PutPolicy{
		Scope: global.CONFIG.Qiniu.BucketName,
	}
	mac := qbox.NewMac(
		global.CONFIG.Qiniu.AccessKey,
		global.CONFIG.Qiniu.SecretKey,
	)
	upToken := putPolicy.UploadToken(mac)

	// 获取存储地区
	zone, err := storage.GetZone(global.CONFIG.Qiniu.AccessKey, global.CONFIG.Qiniu.BucketName)
	if err != nil {
		global.LOG.Error("function storage.GetZone Failed", zap.Any("err", err.Error()))
		return "", errors.New("function storage.GetZone Failed, err:" + err.Error())
	}

	cfg := storage.Config{
		UseHTTPS:      global.CONFIG.Qiniu.UseHTTPS,
		UseCdnDomains: global.CONFIG.Qiniu.UseCdnDomains,
		Zone:          zone,
	}
	ret := storage.PutRet{}
	// 构建表单上传的对象
	formUploader := storage.NewFormUploader(&cfg)
	//额外可选配置
	putExtra := storage.PutExtra{}

	//生成随机文件名
	fileName := utils.RandomFileName(fileHeader.Filename, path...)
	err = formUploader.Put(context.Background(), &ret, upToken, fileName, file, fileHeader.Size, &putExtra)
	if err != nil {
		global.LOG.Error("function formUploader.Put() Failed", zap.Any("err", err.Error()))
		return "", errors.New("function formUploader.Put() Failed, err:" + err.Error())
	}

	return global.CONFIG.Qiniu.CdnPath + "/" + fileName, nil
}

func (*QiNiu) DeleteFile(key string) error {

	key = strings.Replace(key, global.CONFIG.Qiniu.CdnPath+"/", "", 1)
	mac := qbox.NewMac(
		global.CONFIG.Qiniu.AccessKey,
		global.CONFIG.Qiniu.SecretKey,
	)

	// 获取存储地区
	zone, err := storage.GetZone(global.CONFIG.Qiniu.AccessKey, global.CONFIG.Qiniu.BucketName)
	if err != nil {
		global.LOG.Error("function storage.GetZone Failed", zap.Any("err", err.Error()))
		return errors.New("function storage.GetZone Failed, err:" + err.Error())
	}

	cfg := storage.Config{
		UseHTTPS:      global.CONFIG.Qiniu.UseHTTPS,
		UseCdnDomains: global.CONFIG.Qiniu.UseCdnDomains,
		Zone:          zone,
	}
	bucketManager := storage.NewBucketManager(mac, &cfg)
	err = bucketManager.Delete(global.CONFIG.Qiniu.BucketName, key)
	if err != nil {
		global.LOG.Error("function bucketManager.Delete() Failed", zap.Any("err", err.Error()))
		return errors.New("function bucketManager.Delete() Failed, err:" + err.Error())
	}
	return nil
}
