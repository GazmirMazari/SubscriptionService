package main

import (
	"database/sql"
	"log"
	"os"
	"time"

	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
)

const webPort = "80"

func main() {
	//connect to the database
	db := initDb()
	db.Ping()
	//create session

	//create channels

	//create waitgroups

	//set up the application config

	//set up mail

	//listen for web connections

}

func initDb() *sql {
	//connect to the database
	conn := connectToDb()
	if conn == nil {
		log.Panic("can't connect to the database")
	}

	return conn
}

func connectToDb() *sql.DB {
	counts := 0

	dsn := os.Getenv("DB_DSN")

	for {
		connection, err := OpenDB(dsn)
		if err != nil {
			log.Println("can't connect to the database")
		} else {
			log.Println("connected to the database")
			return connection
		}

		if counts > 10 {
			return nil
		}

		log.Print("Backing off for 1 second")
		time.Sleep(1 * time.Second)
		counts++
	}

}

func OpenDB(dsn string) (*sql.DB, error) {

	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
