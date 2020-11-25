package model

import (
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx"
	"github.com/joho/godotenv"
)

var db *pgx.ConnPool

func InitDB() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Println("Error loading .env file")
	}

	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")

	connString := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=disable", dbUser, dbPass, dbHost, dbPort, dbName)
	pgxConf, err := pgx.ParseConnectionString(connString)

	if err != nil {
		log.Println(err)
		return
	}

	pgxPollConf := pgx.ConnPoolConfig{
		ConnConfig:     pgxConf,
		MaxConnections: 5,
	}

	db, err = pgx.NewConnPool(pgxPollConf)
	if err != nil {
		log.Println(err)
		return
	}
}
