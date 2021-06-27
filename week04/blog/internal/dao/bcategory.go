package dao

import "gorm.io/gorm"

type Bcategory struct {
	gorm.Model
	ID      uint   `json:"id" uri:"id"`
	Pid     uint32 `json:"pid" uri:"pid"`
	CatName string `json:"cat_name" uri:"cat_name"`
}

func (bc *Bcategory) Create() {
	Blogdb.Create(bc)
}

func (bc *Bcategory) FindOne(id interface{}) {
	Blogdb.Find(bc, id)
}

func (bc *Bcategory) List() (results []Bcategory) {
	Blogdb.Find(&results)
	return
}

