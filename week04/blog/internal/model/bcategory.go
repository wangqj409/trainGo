package model

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Bcategory struct {
	gorm.Model
	Id       uint32    `json:"id" uri:"id"`
	Pid      uint32    `json:"pid" uri:"pid"`
	CatName  string    `json:"cat_name" uri:"cat_name"`
	CreateAt time.Time `json:"create_at"`
	DeleteAt time.Time `json:"delete_at"`
}

func (c *Bcategory) List() []Bcategory {
	var results []Bcategory

	return results
}
