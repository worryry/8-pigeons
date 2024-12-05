package mysql

import (
	"fmt"
	"github.com/worryry/8-pigeons/pkg/setting"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"time"
)

var DB *gorm.DB

func Start() {
	isEnabled := setting.GetBool("mysql.enable")
	if isEnabled {
		DB = DbInit()
	} else {
		log.Println("不启用mysql")
	}
}

func DbInit() *gorm.DB {
	dbName := setting.GetString("mysql.dbName")
	user := setting.GetString("mysql.user")
	pwd := setting.GetString("mysql.pwd")
	prefix := setting.GetString("mysql.tablePrefix")
	host := setting.GetString("mysql.host")

	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=True&loc=Local",
		user,
		pwd,
		host,
		dbName)
	ClientDb, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   prefix,
			SingularTable: true, //开启表明复数形式
		},
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Println(err)
	}
	sqlDB, err := ClientDb.DB()
	// 最大空闲数
	sqlDB.SetMaxIdleConns(setting.GetInt("mysql.maxIdleConns"))
	// 最大连接数
	sqlDB.SetMaxOpenConns(setting.GetInt("mysql.maxOpenConns"))
	// 最长连接时间
	sqlDB.SetConnMaxLifetime(5 * time.Minute)
	return ClientDb
}
