package config

import (
	"context"
	"github.com/jackc/pgx/v4"
	"github.com/joho/godotenv"
	"log"
	"os"
)

var (
	host = ""
	port = ""
	user = ""
	password = ""
	dbname = ""
	dbUrl = ""
)

var (
	DB *pgx.Conn
	DBerr error
)

func init() {

	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	} else
	{
		host = os.Getenv("PG_HOST")
		port = os.Getenv("PG_PORT")
		user = os.Getenv("PG_USER")
		password = os.Getenv("PG_PASSWORD")
		dbname = os.Getenv("PG_DB")
		dbUrl = "postgres://" + user + ":" + password + "@" + host + ":" + port + "/" + dbname

		log.Println("Database: Connecting to database at:", dbUrl)

		DB, DBerr = pgx.Connect(context.Background(), dbUrl)

		if DBerr != nil {
			log.Fatal(DBerr)
		} else {
			log.Println("Database: Successfully connected!")
		}
	}

}


// Ping test function for the package database
func TestPing() {

	// Pings the global database
	pingError := DB.Ping(context.Background())

	if pingError != nil {
		// An error was returned while pinging the database
		log.Fatal(pingError)
	} else {
		// Database Ping successful
		log.Println("Database: Ping successful.")
	}

}
