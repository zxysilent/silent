package model

import (
	"log"

	_ "github.com/go-sql-driver/mysql" // init
	"github.com/jmoiron/sqlx"
)

var Db *sqlx.DB

func init() {
	db, err := sqlx.Open(`mysql`, `root:root@tcp(127.0.0.1:3306)/news?charset=utf8&parseTime=true`)
	if err != nil {
		log.Fatalln(err.Error())
	}
	Db = db
}
