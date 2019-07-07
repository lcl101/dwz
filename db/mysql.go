package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db *gorm.DB
)

type FrUrl struct {
	// Id      uint   `json:"id"`
	Key     string `json:"key"`
	Url     string `json:"url"`
	Expires int    `json:"expires"`
	gorm.Model
}

func init() {
	//open a db connection
	var err error
	db, err = gorm.Open("mysql", "root:12345@/demo?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect database")
	}
	//Migrate the schema
	db.AutoMigrate(&FrUrl{})
}

func Save(url *FrUrl) {
	db.Save(url)
}

func Get(key string) FrUrl {
	var url FrUrl
	db.First(&url, "key=?", key)
	return url
}
