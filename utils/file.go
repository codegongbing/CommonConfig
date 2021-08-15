package utils

import (
	"CommonConfig/global"
	"fmt"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap/zapcore"
	"math/rand"
	"os"
	"path"
	"strconv"
	"strings"
	"time"
	"unsafe"
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
		"local":      false,
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
	fmt.Println(prefixDir)
	currentTime := time.Now().UnixNano()
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	var src = rand.NewSource(time.Now().UnixNano())
	b := make(
		[]byte,
		10,
	)
	for i, cache, remain := 10-1, src.Int63(), letterIdMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdMax
		}
		if idx := int(cache & letterIdMask); idx < len(letters) {
			b[i] = letters[idx]
			i--
		}
		cache >>= letterIdBits
		remain--
	}
	var str strings.Builder
	str.WriteString(strconv.FormatInt(currentTime, 10))
	str.WriteString(*(*string)(unsafe.Pointer(&b)))
	str.WriteString(path.Ext(fileName))
	return path.Join(prefixDir,str.String())
}
