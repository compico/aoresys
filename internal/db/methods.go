package db

import (
	"context"
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/compico/aoresys/internal/userutil"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
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
		newuuid, err := GenerateUUID(usr.Model)
		if err != nil {
			return err
		}
		_, err = col.InsertOne(sc, bson.M{
			"_id":         newuuid,
			"username":    strings.ToLower(usr.Username),
			"nick":        usr.Username,
			"email":       usr.Email,
			"model":       usr.Model,
			"registrDate": time.Now().Round(time.Second),
			"lastlogin":   time.Now().Round(time.Second),
			"password":    usr.Password,
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

func (db *DB) GetLogin(ctx context.Context, username string, password string, rmbr bool) (*http.Cookie, error) {
	col := db.Client.Database("golosovanie").Collection("users")
	cookiesCol := db.Client.Database("golosovanie").Collection("cookies")
	err := db.Client.Ping(ctx, readpref.Primary())
	if err != nil {
		return nil, err
	}
	q := bson.M{"username": username}
	opts := options.FindOne().SetProjection(bson.M{"username": 1, "password": 1, "_id": 1})
	res := col.FindOne(ctx, q, opts)
	if err != nil {
		return nil, err
	}
	var result struct {
		Username string
		Password string
		_ID      string
	}
	err = res.Decode(&result)
	if err != nil && !errors.Is(err, mongo.ErrNoDocuments) {
		return nil, err
	}
	if errors.Is(err, mongo.ErrNoDocuments) {
		return nil, UserNotCreatedError
	}
	if result.Password != password {
		return nil, WrongPassError
	}
	if result.Username == username && result.Password == password {
		q = bson.M{"_id": result._ID}
		res = cookiesCol.FindOne(ctx, q)
		if err != nil && !errors.Is(err, mongo.ErrNoDocuments) {
			return nil, err
		}
		if errors.Is(err, mongo.ErrNoDocuments) {
			cookie := &http.Cookie{
				Name:     "_SECURE_LOGIN",
				Value:    result._ID,
				Secure:   true,
				HttpOnly: true,
				SameSite: http.SameSiteNoneMode,
			}
			cookie.Expires = time.Now().AddDate(0, 0, 1)
			if rmbr {
				cookie.Expires = time.Now().AddDate(1, 0, 0)
			}
			session, err := db.Client.StartSession()
			if err != nil {
				return nil, err
			}
			err = session.StartTransaction()
			if err != nil {
				return nil, err
			}
			err = mongo.WithSession(ctx, session, func(sc mongo.SessionContext) error {
				_, err = cookiesCol.InsertOne(sc, bson.M{
					"_id":        result._ID,
					"username":   username,
					"cookie":     cookie,
					"createTime": time.Now().Round(time.Second),
					"expireTime": cookie.Expires,
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
				return nil, err
			}
			return cookie, nil
		}
		var cookieResult struct {
			_ID        string
			Username   string
			Cookie     http.Cookie
			CreateTime time.Time
			ExpireTime time.Time
		}
		err = res.Decode(&cookieResult)
		if err != nil {
			return nil, err
		}
		if cookieResult.ExpireTime.Unix() < time.Now().Unix() {
			//can be problem there :)
			return &cookieResult.Cookie, nil
		}
		q = bson.M{"_id": cookieResult._ID}
		_, err := cookiesCol.DeleteOne(ctx, q)
		if err != nil {
			return nil, err
		}
		db.GetLogin(ctx, username, password, rmbr)
	}
	return nil, nil
}

func (db *DB) ExistUsername(ctx context.Context, username string) (bool, error) {
	col := db.Client.Database("golosovanie").Collection("users")
	err := db.Client.Ping(ctx, readpref.Primary())
	if err != nil {
		return true, err
	}
	validator := userutil.NewValidator(username)
	username = validator.GetRightVar()
	if username == "" {
		return true, errors.New("Неверная длинна или запрещённые символы!")
	}
	q := bson.M{"username": username}
	opts := options.FindOne().SetProjection(bson.M{"username": 1, "_id": 0})
	res := col.FindOne(ctx, q, opts)
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
