package db

import (
	"fmt"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var MySQLDB *gorm.DB
var err error

// @brief 初始化操作
func MysqlCoreInit(username, pwd, addr, port, dbname string) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", username, pwd, addr, port, dbname)
	MySQLDB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger:                 logger.Default.LogMode(logger.Info),
		SkipDefaultTransaction: true,
	})

	if err != nil {
		fmt.Println("无法连接数据库,请查看数据库用户密码或权限等是否有误!")
		os.Exit(1)
	}

	// 获取底层的*sql.DB对象
	sqlDB, err := MySQLDB.DB()
	if err != nil {
		fmt.Println("获取数据库连接对象失败")
		os.Exit(1)
	}

	// TODO: 配置连接池参数(之后可以暴露出去)

	// 设置最大闲置连接数为10
	sqlDB.SetMaxIdleConns(10)

	// 设置最大打开连接数为100
	sqlDB.SetMaxOpenConns(100)

	// 设置连接的最大生命周期是3分钟
	sqlDB.SetConnMaxLifetime(3 * time.Minute)
}
