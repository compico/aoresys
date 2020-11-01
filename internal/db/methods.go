package db

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func (db *DB) AddNewUser(ctx context.Context) error {
	err := db.Client.Connect(ctx)
	if err != nil {
		return err
	}
	defer func() {
		err := db.Client.Disconnect(ctx)
		if err != nil {
			fmt.Printf("Error: %v", err.Error())
		}
	}()
	col := db.Client.Database("golosovanie").Collection("users")
	session, err := db.Client.StartSession()
	if err != nil {
		return err
	}
	err = session.StartTransaction()
	if err != nil {
		return err
	}
	err = mongo.WithSession(ctx, session, func(sc mongo.SessionContext) error {
		_, err = col.InsertOne(sc, bson.M{
			"user":     "test",
			"password": "test2",
		})
		err = session.CommitTransaction(sc)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}
