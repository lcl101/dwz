package slot

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/lcl101/dwz/conf"
)

const (
	RedirectProxy = "proxy"
	Redirect301   = "301"
	Redirect302   = "302"
	Redirect303   = "303"
	Redirect307   = "307"
)

var (
	db *gorm.DB
)

type Url struct {
	Slot         string    `json:"slot" gorm:"primary_key"`
	Origin       string    `json:"origin"`
	RedirectType string    `json:"redirect_type"`
	CreatedTime  time.Time `json:"created_time" `
	ExpiresIn    int       `json:"expires_in"`
	Count        int       `json:"count"`
	Url          string    `json:"url" `
}

type Uri struct {
	AppKey    string `json:"appKey"`
	AppSecret string `json:"appSecret"`
	Url       string `json:"url"`
}

func InitDB() {
	//open a db connection
	var err error
	dbUrl := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		conf.Conf.Db.User, conf.Conf.Db.Passwd, conf.Conf.Db.Host, conf.Conf.Db.Port, conf.Conf.Db.Database)
	db, err = gorm.Open("mysql", dbUrl)
	if err != nil {
		panic("failed to connect database")
	}
	//Migrate the schema
	db.AutoMigrate(&Url{})
}

func NewUrl(origin string, unique bool) *Url {
	//check origin exist
	if unique && OriginExist(origin) {
		url := FindUrlByOrigin(origin)
		return url
	}
	url := &Url{
		Slot:         SlotGenerator.Get(),
		Origin:       origin,
		RedirectType: Redirect301,
		CreatedTime:  time.Now(),
		ExpiresIn:    0,
		Count:        0,
	}
	for {
		// check slot exist
		if !SlotExist(url.Slot) {
			break
		}
		url.Slot = SlotGenerator.Get()
	}
	db.Save(url)
	return url
}

// func NewCustomUrl(slot, origin string) *Url {
// 	if SlotExist(slot) {
// 		return nil
// 	}
// 	url := &Url{
// 		Slot:         slot,
// 		Origin:       origin,
// 		RedirectType: Redirect301,
// 		CreatedTime:  time.Now(),
// 		ExpiresIn:    0,
// 		Count:        0,
// 	}
// 	url.Save()
// 	return url
// }

func FindUrlByOrigin(origin string) *Url {
	var url Url
	db.First(&url, "origin=?", origin)
	return &url
}

func FindUrlBySlot(slot string) *Url {
	// session := GetMgo()
	// defer session.Close()
	// var url Url
	// err := session.DB("ifth").C("url").Find(bson.M{"slot": slot}).One(&url)
	// if err != nil {
	// 	return nil, err
	// }
	// url.Count++
	// url.Save()
	var url Url
	db.First(&url, slot)
	return &url
}

// func FindHottestUrls(limit int) ([]Url, error) {
// 	session := GetMgo()
// 	defer session.Close()
// 	var urls []Url
// 	err := session.DB("ifth").C("url").Find(bson.M{}).Sort("-count").Limit(limit).All(&urls)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return urls, nil
// }

// func FindNewestUrls(limit int) ([]Url, error) {
// 	session := GetMgo()
// 	defer session.Close()
// 	var urls []Url
// 	err := session.DB("ifth").C("url").Find(bson.M{}).Sort("-_id").Limit(limit).All(&urls)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return urls, nil
// }

func OriginExist(origin string) bool {
	// session := GetMgo()
	// defer session.Close()
	// ct, err := session.DB("ifth").C("url").Find(bson.M{"origin": origin}).Count()
	// if err != nil {
	// 	return false //may cause problem?
	// }
	// if ct > 0 {
	// 	return true
	// }
	// return false
	url := FindUrlByOrigin(origin)
	return url == nil
}

func SlotExist(slot string) bool {
	// session := GetMgo()
	// defer session.Close()
	// ct, err := session.DB("ifth").C("url").Find(bson.M{"slot": slot}).Count()
	// if err != nil {
	// 	return false //may cause problem?
	// }
	// if ct > 0 {
	// 	return true
	// }
	// return false
	url := FindUrlBySlot(slot)
	return url == nil
}

func (u *Url) Expired() bool {
	if u.ExpiresIn > 0 {
		if int(time.Now().Sub(u.CreatedTime).Seconds()) > u.ExpiresIn {
			return true
		}
	}
	return false
}

// func (u *Url) Save() error {
// 	session := GetMgo()
// 	defer session.Close()
// 	_, err := session.DB("ifth").C("url").Upsert(bson.M{"slot": u.Slot}, u)
// 	return err
// }
