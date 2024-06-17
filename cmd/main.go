package main

import (
	"Ecom/cmd/api"
	"Ecom/cmd/db"
	"Ecom/config"
	"database/sql"
	"log"
	"github.com/go-sql-driver/mysql"
)


func main(){

	// create a new database with the provided credentials
	db, err := db.NewMySQLStorage(mysql.Config{
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

	//Ensures that current application is connected to the database and if
	// connects it to the database 
	initStorage(db)

	// sends the database that is ready to The APIServer constructor to create a new 
	// a server with the port number 
	server := api.NewAPIServer(":8080", db)

	//Once the server is ready, we start the server the server by calling its run method
	if error := server.Run(); error != nil {
		log.Fatal(error)
	}
}

func initStorage(db *sql.DB){
	err := db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("DB: Successfully connected")
}