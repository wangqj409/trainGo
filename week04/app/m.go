package main

import (
	_ "github.com/wangqj409/trainGo/week04/app/internal/model"
	"github.com/wangqj409/trainGo/week04/app/internal/model"
	"fmt"
	"encoding/json"
	"log"
)

type val struct {
	Name  string
	Short string
}

func main() {
	//model.Mdb.Begin()
	//bs := []dao.Brand{
	//	{Name: "三星", ShortName: "sunsamg"},
	//	{Name: "诺基亚", ShortName: "Nokia"},
	//	{Name: "苹果", ShortName: "apple"},
	//}
	//model.Mdb.Create(&bs)
	//model.Mdb.Commit()
	//fmt.Printf("%v\n", bs)
	b := model.Brand{}
	res := b.List()

	c := model.Category{}
	allc := c.List()
	cs , _ := json.Marshal(&allc)
	fmt.Printf("%s\n", cs)

	s, err := json.Marshal(&res)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", string(s))

}
