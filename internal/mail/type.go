package mail

import "net/smtp"

type Client struct {
	Auth   smtp.Auth
	Config Config
}

type Message struct {
	to      []string
	message string
	subject string
	Pack    Pack
	Packs   Packs
}

type Pack []byte
type Packs []Pack

type Config struct {
	Username string `json:"user"`
	Password string `json:"pass"`
	Hostname string `json:"host"`
	Port     string `json:"port"`
	Email    string `json:"email"`
}
