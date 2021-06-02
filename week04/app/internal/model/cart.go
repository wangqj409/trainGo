package model

import "github.com/wangqj409/trainGo/week04/app/internal/dao"

type Cart struct {
	dao.Cart
}

func (b *Cart) List() []dao.Cart {
	var carts []dao.Cart
	Mdb.Select("goods_id, goods_name, goods_price").Find(&carts, "id > ?", 0)
	return carts
}
