package main

import (
	// "Ecom/cmd/api"
	"Ecom/cmd/db"
	"Ecom/config"
	"os"

	// "database/sql"
	"log"

	mysqlConfig "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main(){
	
	// create a new database with the provided credentials
	db, err := db.NewMySQLStorage(mysqlConfig.Config{
		User: config.Envs.DBUser,
		Passwd: config.Envs.DBPassword,
		Addr: config.Envs.DBAddress,
		DBName: config.Envs.DBName,
		Net: "tcp",
		AllowNativePasswords: true,
		ParseTime: true,
	})

	if err != nil {
		log.Fatal(err)
	}

	driver, error := mysql.WithInstance(db, &mysql.Config{})
	if error != nil {
		log.Fatal(error)
	}

	m, error := migrate.NewWithDatabaseInstance(
		"file://cmd/migrate/migrations",
		"mysql",
		driver ,
	)
	if error != nil {
		log.Fatal(error)
	}

	cmd := os.Args[(len(os.Args) - 1)]

	if cmd == "up"{
		//if there was an error during the migration. 
		//If err is not nil and it's not migrate.ErrNoChange, it logs the error
		if err := m.Up(); err != nil && err != migrate.ErrNoChange{
			log.Fatal(err)
		}
	}
	if cmd == "down"{
		if err := m.Down(); err != nil && err != migrate.ErrNoChange{
			log.Fatal(err)
		}
	}
}