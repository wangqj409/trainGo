package model

import (
	"gorm.io/gorm"
	_ "gorm.io/gorm/mysql"
)

var (
	db *gorm.DB
)

func init() {
	db = gorm.Open("mysql", &gorm.Config{})
}