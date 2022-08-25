package main

import (
	"context"
	"log"
	"os"

	"github.com/go-pg/pg/v10"
)

var db *pg.DB

func connectToDB() {
	db = pg.Connect(&pg.Options{
		Addr:     os.Getenv("DB_ADDR"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASS"),
		Database: os.Getenv("DB_NAME"),
	})
	ctx := context.Background()
	if err := db.Ping(ctx); err != nil {
		log.Fatalln(err)
	}
}
