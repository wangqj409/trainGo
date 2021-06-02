package dao

import "gorm.io/gorm"

type Brand struct {
	gorm.Model
	Name      string `json:"name"`
	ShortName string `json:"short"`
}

//func (b Brand) TableName() string {
//	return "brands"
//}
