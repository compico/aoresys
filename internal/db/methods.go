package db

import (
	"context"
	"errors"
	"strings"

	"github.com/compico/aoresys/internal/userutil"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"golang.org/x/crypto/bcrypt"
)

func (db *DB) AddNewUser(ctx context.Context, usr userutil.User) error {
	err := db.Client.Ping(ctx, readpref.Primary())
	if err != nil {
		return err
	}
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
		newuuid, err := generateUUID(usr.Model)
		if err != nil {
			return err
		}
		hashpass, err := bcrypt.GenerateFromPassword([]byte(usr.Password), bcrypt.DefaultCost)
		if err != nil {
			return err
		}
		_, err = col.InsertOne(sc, bson.M{
			"_id":      newuuid,
			"username": strings.ToLower(usr.Username),
			"nick":     usr.Username,
			"password": string(hashpass),
		})
		if err != nil {
			return err
		}
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

func (db *DB) ExistUsername(ctx context.Context, username string) (bool, error) {
	col := db.Client.Database("golosovanie").Collection("users")
	err := db.Client.Ping(ctx, readpref.Primary())
	if err != nil {
		return true, err
	}
	opts := options.FindOne().SetProjection(bson.M{"username": 1, "_id": 0})
	res := col.FindOne(ctx, bson.M{"username": username}, opts)
	if err != nil {
		return true, err
	}
	var result struct {
		Username string
	}
	err = res.Decode(&result)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return false, nil
		}
		return true, err
	}
	if result.Username == username {
		return true, nil
	}
	return false, nil
}
