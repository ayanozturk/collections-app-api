package database

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"log"
)

var DB *sqlx.DB

func InitDB(dataSourceName string) {
	var err error
	DB, err = sqlx.Connect("mysql", dataSourceName)
	if err != nil {
		log.Fatalln(err)
	}

	DB.SetMaxOpenConns(25)
	DB.SetMaxIdleConns(25)
	DB.SetConnMaxLifetime(5)
}
