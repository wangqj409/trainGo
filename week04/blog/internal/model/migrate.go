package model

import (
	"gorm.io/gorm"
	"gorm.io/driver/mysql"
)

var (
	blogdb *gorm.DB
)

func init() {
	dsn := "root:@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	blogdb, _ = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	blogdb.AutoMigrate(&Bcategory{})
}
