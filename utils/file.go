package utils

import (
	"CommonConfig/global"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap/zapcore"
	"os"
	"path"
	"strconv"
	"time"
)

const (
	// 6 bits to represent a letter index
	letterIdBits = 6
	// All 1-bits as many as letterIdBits
	letterIdMask = 1<<letterIdBits - 1
	letterIdMax  = 63 / letterIdBits
)

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func GetWriteSyncer() zapcore.WriteSyncer {
	now := time.Now()
	fileName := fmt.Sprintf("log/%d-%d-%d.log", now.Year(), now.Month(), now.Day())

	lumberJackLogger := &lumberjack.Logger{
		Filename:   fileName,
		MaxSize:    10,
		MaxBackups: 5,
		MaxAge:     30,
		Compress:   false,
	}
	return zapcore.AddSync(lumberJackLogger)
}

func RandomFileName(fileName string, Path ...string) string {
	var datePath = map[string]bool{
		"qiniu":      global.CONFIG.Qiniu.DatePath,
		"local":      global.CONFIG.Local.DatePath,
		"aliyun-oss": global.CONFIG.AliyunOSS.DatePath,
		"TencentCOS": false,
	}
	var prefixDir string
	if len(Path) == 0 {
		prefixDir = ""
	} else {
		prefixDir = Path[0]
	}
	if datePath[global.CONFIG.Project.OSSType] {
		now := time.Now()
		year := strconv.Itoa(now.Year())
		month := strconv.Itoa(int(now.Month()))
		day := strconv.Itoa(now.Day())
		prefixDir = path.Join(prefixDir, year, month, day)
	}
	h := sha1.New()
	h.Write([]byte(fileName + strconv.FormatInt(time.Now().UnixNano(), 10)))
	name := hex.EncodeToString(h.Sum(nil))
	return path.Join(prefixDir, name+path.Ext(fileName))
}
