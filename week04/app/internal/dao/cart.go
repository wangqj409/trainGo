package dao

import "gorm.io/gorm"

type Cart struct {
	gorm.Model
	Uid         uint32
	GoodsId     uint32
	GoodsName   string
	MarketName  string
	GoodsPrice  float64
	Promotion   string
	GoodsNum    uint32
	GoodsSN     string
	GoodsAttrId string
	IsGift      uint8
	IsReal      uint8
}
