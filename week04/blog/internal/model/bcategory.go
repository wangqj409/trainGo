package model

import (
	"time"
)

type Bcategory struct {
	Id       uint32    `json:"id" uri:"id"`
	Pid      uint32    `json:"pid" uri:"pid"`
	CatName  string    `json:"cat_name" uri:"cat_name"`
	CreateAt time.Time `json:"create_at"`
	DeleteAt time.Time `json:"delete_at"`
}

func (c *Bcategory) List() []Bcategory {
	var results []Bcategory
	blogdb.Find(&results)
	return results
}

func (c *Bcategory) One(id uint32) Bcategory {
	var result Bcategory

	blogdb.Find(&result, Bcategory{
		Id: id,
	})
	return result
}

func (c *Bcategory) Create() *Bcategory {
	c.CreateAt = time.Now()
	blogdb.Select("pid,cat_name,create_at").Create(c)
	if blogdb.RowsAffected > 0 {
		return c
	}
	return &Bcategory{}
}
