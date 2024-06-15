package db

import (
	"database/sql"
	"log"
	"github.com/go-sql-driver/mysql"
)

func NewMySQLStorage(cfg mysql.Config)(*sql.DB, error){

	db, error := sql.Open("mysql", cfg.FormatDSN())

	if error != nil{
		log.Fatal(error)
	}
	return db, nil
}

