package db

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

// init the database connection
// param:
// - dsn: data source name
func InitDB(dsn string) error {
	var err error

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Printf("数据库连接失败: %v", err)
		return err
	}

	sqlDB, err := DB.DB()
	if err != nil {
		return err
	}
	
	sqlDB.SetMaxIdleConns(10)  // 空闲连接池最大数量
	sqlDB.SetMaxOpenConns(100)  // 打开数据库连接的最大数量

	log.Printf("数据库连接成功！")
	return nil
}