package mail

import "net/smtp"

func NewClient() *Client {
	return new(Client)
}
func (client *Client) NewAuth(cfg Config) {
	client.Auth = smtp.PlainAuth("", cfg.Username, cfg.Password, cfg.Hostname)
}

func (client *Client) SendMail(msg Message) error {
	if len(msg.Packs) == 0 {
		err := smtp.SendMail(
			"smtp.mailtrap.io:25",
			client.Auth,
			"dc791412df-af4ac5@inbox.mailtrap.io",
			msg.to,
			msg.Pack,
		)
		if err != nil {
			return err
		}
		return nil
	}
	if i := len(msg.Packs); i > 1 {
		for l := 0; l < i; l++ {
			err := smtp.SendMail(
				"smtp.mailtrap.io:25",
				client.Auth,
				"dc791412df-af4ac5@inbox.mailtrap.io",
				msg.to,
				msg.Packs[l],
			)
			if err != nil {
				return err
			}
		}
		return nil
	}
	return nil
}
