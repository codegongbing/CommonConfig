package upload

import (
	"CommonConfig/global"
	"CommonConfig/utils"
	"errors"
	"go.uber.org/zap"
	"io"
	"mime/multipart"
	"os"
	"strings"
)

type Local struct{}

func (*Local) UploadFile(fileHeader *multipart.FileHeader, path ...string) (string, error) {
	file, err := fileHeader.Open()
	if err != nil {
		global.LOG.Error("function fileHeader.Open() Failed", zap.Any("err", err.Error()))
		return "", errors.New("function fileHeader.Open() Failed, err:" + err.Error())
	}
	defer func(file multipart.File) {
		err = file.Close()
		if err != nil {
			global.LOG.Error("function file.Close() Failed", zap.Any("err", err.Error()))
		}
	}(file)
	name := utils.RandomFileName(fileHeader.Filename, path...)

	filePath := global.CONFIG.Local.Path + "/" + name
	split := strings.Split(filePath, "/")
	dir := strings.Replace(filePath, "/"+split[len(split)-1], "", 1)

	err = os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		global.LOG.Error("os.MkdirAll() Failed", zap.Any("err", err.Error()))
		return "", errors.New("os.MkdirAll() Failed, err:" + err.Error())
	}

	out, err := os.Create(filePath)
	if err != nil {
		global.LOG.Error("os.Create() Failed", zap.Any("err", err.Error()))

		return "", errors.New("os.Create() Failed, err:" + err.Error())
	}
	defer func(out *os.File) {
		err := out.Close()
		if err != nil {
			global.LOG.Error("out.Close() Failed", zap.Any("err", err.Error()))
		}
	}(out) // 创建文件 defer 关闭

	_, err = io.Copy(out, file) // 传输（拷贝）文件
	if err != nil {
		global.LOG.Error("io.Copy() Failed", zap.Any("err", err.Error()))
		return "", errors.New("io.Copy() Failed, err:" + err.Error())
	}
	return filePath, nil
}

func (*Local) DeleteFile(key string) error {

	return nil
}
