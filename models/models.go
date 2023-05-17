package models

import (
	"log"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var (
	db   *sqlx.DB
	Time = time.Now()
)

func GetDB() *sqlx.DB {
	return db
}

func SetupDB() {

	var err error
	// DB="root:1365@/kondor?parseTime=true" go run main.go
	db, err = sqlx.Open("mysql", os.Getenv("DB"))
	if err != nil {

		log.Fatal(err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

}
