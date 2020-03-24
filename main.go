package main

import (
	"context"
	"github.com/jackc/pgx"
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

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	} else {
		host		= os.Getenv("PG_HOST")
		port		= os.Getenv("PG_PORT")
		user		= os.Getenv("PG_USER")
		password 	= os.Getenv("PG_PASSWORD")
		dbname 		= os.Getenv("PG_DB")
		dbUrl 		= "postgres://"+user+":"+password+"@"+host+":"+port+"/"+dbname
	}
}

var (
	conn *pgx.Conn
	err error
)

func main() {

	conn, err = pgx.Connect(context.Background(), dbUrl)
	if err != nil {
		panic(err)
	}

	defer conn.Close(context.Background())

	err = conn.Ping(context.Background())

	if err != nil {
		panic(err)
	} else {
		log.Println("Successfully connected to database!")
	}

}
