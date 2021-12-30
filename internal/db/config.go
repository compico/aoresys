package db

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/url"
)

func NewConfig(path string) (*MongoConfig, error) {
	mongodb := new(MongoConfig)
	err := mongodb.readConfig(path)
	if err != nil {
		return nil, err
	}
	err = mongodb.getUri()
	if err != nil {
		return nil, err
	}
	return mongodb, nil
}

func (mongodb *MongoConfig) getUri() error {
	//"mongodb://<username>:<password>@<cluster-address>/test?w=majority"
	var user *url.Userinfo
	if mongodb.User != "" {
		if mongodb.Password != "" {
			user = url.UserPassword(mongodb.User, mongodb.Password)
		}
		if mongodb.Password == "" {
			user = url.User(mongodb.User)
		}
	}
	if len(mongodb.Hostname) == 0 {
		return errors.New("Config error! Field 'HostName' is nil")
	}
	query := "retryWrites=true&w=majority"
	x := url.URL{
		Scheme:   "mongodb+srv",
		User:     user,
		Host:     mongodb.getHostname(),
		Path:     "/golosovanie",
		RawQuery: query,
	}
	mongodb.URI = x.String()
	return nil
}

func (mongodb *MongoConfig) getHostname() string {
	if len(mongodb.Hostname) == 1 {
		return mongodb.Hostname[0]
	}
	var hosts string
	for i := 0; i < len(mongodb.Hostname); i++ {
		hosts += mongodb.Hostname[i]
		if i == len(mongodb.Hostname)-1 {
			break
		}
		hosts += ","
	}
	return hosts
}

func (config *MongoConfig) readConfig(path string) error {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	err = json.Unmarshal(file, config)
	if err != nil {
		return err
	}
	return nil
}
