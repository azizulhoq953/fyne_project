package main

import (
	"database/sql"
	"log"
)

func dbcon() {
	db, err = sql.Open("sqlite3", "E:/GOLANG/src/master_academy/golang/minierp/mini-erp.db")
	if err != nil {
		log.Fatal(err)
	}
	db.SetMaxOpenConns(1)

	log.Println("db connection successful")

}
