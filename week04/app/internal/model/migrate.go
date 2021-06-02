package model

import (
	"fmt"
	"os"
	"gorm.io/driver/mysql"
	"github.com/wangqj409/trainGo/week04/app/internal/dao"
	"gopkg.in/yaml.v2"
	"gorm.io/gorm"
)

var Mdb *gorm.DB

type dbconf struct {
	Services struct {
		Mysql conf `yaml:"mysql"`
		Redis conf `yaml:"redis"`
	}
}
type conf struct {
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Port     string `yaml:"port"`
	Host     string `yaml:"host"`
	Dbname   string `yaml:"dbname"`
	Charset  string `yaml:"charset"`
}

func init() {
	initdb()
}

func initdb() {
	_mysqlconf := readfile("src/github.com/wangqj409/trainGo/week04/app/config/mysql.yaml")
	mconf := &dbconf{}
	yaml.Unmarshal(_mysqlconf, &mconf)
	dsn := fmt.Sprintf("%s", &(mconf.Services.Mysql))
	//println(dsn)
	Mdb, _ = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	Mdb.AutoMigrate(dao.Brand{}, dao.Category{}, dao.Goods{}, dao.Cart{})
}

func (conf *conf) String() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true",
		conf.User,
		conf.Password,
		conf.Host,
		conf.Port,
		conf.Dbname,
		conf.Charset)
}

func readfile(pathname string) []byte {
	_bs, err := os.ReadFile(pathname)
	if err != nil {
		panic(err)
	}
	return _bs
}
