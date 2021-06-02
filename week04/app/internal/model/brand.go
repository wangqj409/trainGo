package model

import (
	"github.com/wangqj409/trainGo/week04/app/internal/dao"
)

type Brand struct {
	dao.Brand
}

func (b *Brand) List() []dao.Brand {
	var brands []dao.Brand
	Mdb.Select("name, short_name").Find(&brands, "id > ?", 0)
	return brands
}

func (b *Brand) Create() {

}
