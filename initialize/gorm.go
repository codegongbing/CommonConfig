package initialize

import (
	"CommonConfig/global"
	"CommonConfig/initialize/internal"
	"fmt"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var sqlType = map[string]int{
	"mysql":      1,
	"postgresql": 2,
}

func Gorm() *gorm.DB {
	var db *gorm.DB
	switch sqlType[global.CONFIG.Sql.GormConfig.SqlType] {
	case 1:
		db = initMysql()
	case 2:
		db = initPostgres()
	}
	sqlDB, err := db.DB()
	if err != nil {
		global.LOG.Error("db.DB() failed", zap.Any("err", err.Error()))
		return nil
	}
	sqlDB.SetMaxIdleConns(global.CONFIG.Sql.GormConfig.MaxIdleConn)
	sqlDB.SetMaxOpenConns(global.CONFIG.Sql.GormConfig.MaxOpenConn)
	return db
}

func initMysql() *gorm.DB {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		global.CONFIG.Sql.MysqlConfig.Username,
		global.CONFIG.Sql.MysqlConfig.Password,
		global.CONFIG.Sql.MysqlConfig.Hostname,
		global.CONFIG.Sql.MysqlConfig.Port,
		global.CONFIG.Sql.MysqlConfig.Dbname,
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
	return db
}

func initPostgres() *gorm.DB {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia"+
			"/Shanghai",
		global.CONFIG.Sql.PostgresConfig.Hostname,
		global.CONFIG.Sql.PostgresConfig.Username,
		global.CONFIG.Sql.PostgresConfig.Password,
		global.CONFIG.Sql.PostgresConfig.Dbname,
		global.CONFIG.Sql.PostgresConfig.Port,
	)
	db, err := gorm.Open(
		postgres.Open(dsn),
		gormConfig(),
	)
	if err != nil {
		return nil
	}
	return db
}

func gormConfig() *gorm.Config {
	personalConfig := global.CONFIG.Sql.GormConfig
	config := &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
		SkipDefaultTransaction:                   personalConfig.SkipDefaultTransaction,
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   personalConfig.TablePrefix,
			SingularTable: personalConfig.SingularTable,
		},
	}
	switch global.CONFIG.Sql.GormConfig.LogMode {
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
