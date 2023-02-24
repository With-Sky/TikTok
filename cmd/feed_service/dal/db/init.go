package db

import (
	"fmt"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	gormopentracing "gorm.io/plugin/opentracing"
	"os"
	"tiktok/cmd/comment_service/global"
)

var TIK_DB *gorm.DB

// Init init DB
func Init() {
	v := global.Config.Viper
	if global.Config.Viper.GetString("Mysql.Dbname") == "" {
		global.LOG.Error("获取配置错误")
	}
	dsn := v.GetString("Mysql.Username") + ":" + v.GetString("Mysql.Password") + "@tcp(" + v.GetString("Mysql.Path") + ")/" + v.GetString("Mysql.Dbname") + "?" + v.GetString("Mysql.Config")
	fmt.Println(dsn)
	mysqlConfig := mysql.Config{
		DSN:                       dsn,   // DSN data source name
		DefaultStringSize:         191,   // string 类型字段的默认长度
		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据版本自动配置
	}
	if db, err := gorm.Open(mysql.New(mysqlConfig)); err != nil {
		global.LOG.Error("MySQL启动异常", zap.Any("err", err))
		os.Exit(0)
	} else {
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(v.GetInt("Mysql.MaxIdleConns"))
		sqlDB.SetMaxOpenConns(v.GetInt("Mysql.MaxOpenConns"))
		TIK_DB = db
	}

	if err := TIK_DB.Use(gormopentracing.New()); err != nil {
		global.LOG.Error("数据库链路错误")
		panic(err)
	}
}
