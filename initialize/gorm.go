package initialize

import (
	"CommonConfig/global"
	"CommonConfig/initialize/internal"
	"fmt"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

func Gorm() *gorm.DB {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		global.CONFIG.Mysql.Username,
		global.CONFIG.Mysql.Password,
		global.CONFIG.Mysql.Hostname,
		global.CONFIG.Mysql.Port,
		global.CONFIG.Mysql.Dbname,
	)
	db, err := gorm.Open(
		mysql.New(
			mysql.Config{
				DSN:               dsn,
				DefaultStringSize: 256,
			},
		),
		gormConfig(),
	)
	if err != nil {
		return nil
	}
	sqlDB, err := db.DB()
	if err != nil {
		global.LOG.Error("db.DB() failed", zap.Any("err", err.Error()))
		return nil
	}
	sqlDB.SetMaxIdleConns(global.CONFIG.Mysql.PersonalConfig.MaxIdleConn)
	sqlDB.SetMaxOpenConns(global.CONFIG.Mysql.PersonalConfig.MaxOpenConn)
	return db
}

func gormConfig() *gorm.Config {
	personalConfig := global.CONFIG.Mysql.PersonalConfig
	config := &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
		SkipDefaultTransaction:                   personalConfig.SkipDefaultTransaction,
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   personalConfig.TablePrefix,
			SingularTable: personalConfig.SingularTable,
		},
	}
	switch global.CONFIG.Mysql.PersonalConfig.LogMode {
	case "silent", "Silent":
		config.Logger = internal.Default.LogMode(logger.Silent)
	case "error", "Error":
		config.Logger = internal.Default.LogMode(logger.Error)
	case "warn", "Warn":
		config.Logger = internal.Default.LogMode(logger.Warn)
	case "info", "Info":
		config.Logger = internal.Default.LogMode(logger.Info)
	default:
		config.Logger = internal.Default.LogMode(logger.Info)
	}
	return config
}
