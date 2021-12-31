package conf

import (
	"errors"
	"os"

	"github.com/compico/aoresys/internal/db"
)

var (
	EmptyError       = errors.New("Empty!")
	EnvHostnameError = errors.New("MONGODB_HOSTNAME env is empty!")
	EnvDBnameError   = errors.New("MONGODB_DBNAME env is empty!")
	EnvUsernameError = errors.New("MONGODB_USERNAME env is empty!")
	EnvPassordError  = errors.New("MONGODB_PASSWORD env is empty!")
)

func GetMongoConfigFromEnvironment() (*db.MongoConfig, error) {
	mdbHostname, err := getDataOrError(os.Getenv("MONGODB_HOSTNAME"))
	if err != nil {
		return nil, EnvHostnameError
	}
	mdbDBname, err := getDataOrError(os.Getenv("MONGODB_DBNAME"))
	if err != nil {
		return nil, EnvDBnameError
	}
	mdbUsername, err := getDataOrError(os.Getenv("MONGODB_USERNAME"))
	if err != nil {
		return nil, EnvUsernameError
	}
	mdbPassword, err := getDataOrError(os.Getenv("MONGODB_PASSWORD"))
	if err != nil {
		return nil, EnvPassordError
	}
	cfg := new(db.MongoConfig)
	cfg.Hostname = []string{mdbHostname}
	cfg.DBname = mdbDBname
	cfg.User = mdbUsername
	cfg.Password = mdbPassword
	if err = cfg.GetUri(); err != nil {
		return nil, err
	}
	return cfg, nil
}

func getDataOrError(data string) (string, error) {
	if data == "" {
		return "", EmptyError
	}
	return data, nil
}
