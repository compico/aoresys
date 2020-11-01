package db

import (
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewDB(cfg MongoConfig) (*DB, error) {
	db := new(DB)
	err := db.newClient(cfg)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func (db *DB) newClient(cfg MongoConfig) error {
	var err error
	db.Client, err = mongo.NewClient(options.Client().ApplyURI(cfg.URI))
	if err != nil {
		return err
	}
	return nil
}
