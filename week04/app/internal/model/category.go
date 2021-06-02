package model

import (
	"github.com/wangqj409/trainGo/week04/app/internal/dao"
)

type Category struct {
	dao.Category
}

func (b *Category) List() []dao.Category {
	var categories []dao.Category
	Mdb.Find(&categories, "id > ?", 0)
	return categories
}
