package model

import (
	"fmt"
	"github.com/wangqj409/trainGo/week04/blog/internal/dao"
)

type Bcategory struct {
	dao.Bcategory
}

func (c *Bcategory) String() {
	fmt.Printf("%d %d %s\n", c.ID, c.Pid, c.CatName)
}

func (c *Bcategory) List() []dao.Bcategory {
	return c.Bcategory.List()
}

func (c *Bcategory) FindOne(id interface{}) (result *dao.Bcategory) {
	result = &dao.Bcategory{}
	result.FindOne(id)
	return
}

func (c *Bcategory) Create() *Bcategory {
	c.Bcategory.Create()
	return c
}
