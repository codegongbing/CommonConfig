package utils

import (
	"fmt"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap/zapcore"
	"os"
	"time"
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
