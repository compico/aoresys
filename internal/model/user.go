package model

type User struct {
	Id       int    `pg:"pk_id"`
	Username string `pg:"unique,notnull"`
	Nick     string
	Email    string `pg:"unique,notnull"`
	Password string
	Model    bool
	UUID     string `pg:"unique,notnull"`
}
