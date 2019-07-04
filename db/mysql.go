package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var (
	db *sql.DB
)

func init() {
	conn()
}

func conn() {
	d, err := sql.Open("mysql", "root:5L!r00t-my5q1@tcp(www.southlocal.cn:3306)/gaopin_test")
	if err != nil {
		fmt.Println(err)
	}
	db = d
}

func getUrl(tkey string) string {

}
