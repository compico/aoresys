package db

import "errors"

var (
	UserNotCreatedError = errors.New("Неправильный логин или пароль!")
	WrongPassError      = errors.New("Неправильный логин или пароль!")
)
