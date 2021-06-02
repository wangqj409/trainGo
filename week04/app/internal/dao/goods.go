package dao

import "gorm.io/gorm"

type Goods struct {
	gorm.Model
	Name        string  `json:"good_name"`
	ShopPrice   float64 `json:"goods_price"`
	MarketPrice float64
	Stock       uint32
	Sold        uint32
	Img         string `json:"img"`
}
