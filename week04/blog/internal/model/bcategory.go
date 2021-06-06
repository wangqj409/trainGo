package model

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Bcategory struct {
	gorm.Model
	Id       uint32    `json:"id"`
	Pid      uint32    `json:"pid"`
	CatName  string    `json:"cat_name"`
	CreateAt time.Time `json:"create_at"`
	DeleteAt time.Time `json:"delete_at"`
}

func (c *Bcategory) List() []Bcategory {
	var results []Bcategory

	return results
}
