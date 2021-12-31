package main

import (
	"github.com/compico/aoresys/conf"
	"github.com/compico/aoresys/internal/db"
)

var (
	cdb *db.DB
)

func initDBClient() {
	dbcfg, err := conf.GetMongoConfigFromEnvironment()
	if err != nil {
		panic(err)
	}
	cdb, err = db.NewDB(*dbcfg)
	if err != nil {
		panic(err)
	}
}
