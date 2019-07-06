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
	d, err := sql.Open("mysql", "")
	if err != nil {
		fmt.Println(err)
	}
	db = d
}

func getUrl(tkey string) string {

}
