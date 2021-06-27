package model

import (
	"fmt"
	"github.com/wangqj409/trainGo/week04/blog/internal/dao"
)

type Barticle struct {
	dao.Barticle
}

func (c *Barticle) String() {
	fmt.Printf("%d %d %s\n", c.ID, c.Title, c.SubTitle)
}

func (c *Barticle) List(offset int, limit int) []dao.Barticle {
	return c.Barticle.List(offset, limit)
}

func (c *Barticle) FindOne(id string) (result *dao.Barticle) {
	result = &dao.Barticle{}
	result.FindOne(id)
	return
}

func (c *Barticle) Create() *Barticle {
	c.Barticle.Create()
	return c
}
