package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/go-sql-driver/mysql"
)

var (
	db *gorm.DB
)

func Connect() {
	d, err := gorm.Open("mysql","root:@/bookstore?charset=utf8&parseTime=True&loc=Local")
	if err!=nil{
		panic(err)
	}
	db = d
}
func GetDB() *gorm.DB {
	return db
}


