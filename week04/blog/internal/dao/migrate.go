package dao

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	Blogdb *gorm.DB
)

func newDB(dsn string) *gorm.DB {
	db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	return db
}

func init() {
	dsn := "root:@tcp(127.0.0.1:3306)/blog?charset=utf8mb4&parseTime=True&loc=Local"
	Blogdb = newDB(dsn)
	Blogdb.AutoMigrate(&Bcategory{}, &Barticle{})
}
