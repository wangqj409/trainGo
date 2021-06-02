package dao

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	Name string `json:"cat_name"`
	Pid  uint32 `json:"parent_id"`
}
