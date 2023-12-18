package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

const appPort = "3000"

type Config struct{
	DB     *sql.DB
}

func main() {
	dbConnection := connectToDB()
	if dbConnection == nil {
		log.Panic("Cannot connect to DB")
	}

	app := Config{
		DB: dbConnection,
	}

	server := &http.Server{
		Addr: fmt.Sprintf(":%s", appPort),
		Handler: app.routes(), 
	}

	fmt.Printf("Server is listening on port %s", appPort)
	
	err := server.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}

}

const (
  host     = "localhost"
  port     = 6543
  user     = "postgres"
  password = "123123"
  dbname   = "postgres"
)

func connectToDB() *sql.DB {
	dataSourceName := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", dataSourceName)
	if err != nil {
		return nil
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		return nil
	}

	return db;
}