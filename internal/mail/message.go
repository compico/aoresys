package mail

import (
	"errors"
	"fmt"
	"time"
)

func NewMessage() *Message {
	return new(Message)
}

func (msg *Message) AddTo(addresses ...string) {
	msg.to = append(msg.to, addresses...)
}

func (msg *Message) AddSubject(subject string) {
	msg.subject = subject
}

func (msg *Message) AddMessage(textmsg string) {
	msg.message = textmsg
}

func (msg *Message) CompileMail() error {
	if msg.subject == "" {
		return errors.New("Subject is void!")
	}
	if len(msg.to) == 0 {
		return errors.New("No have email to send!")
	}
	if msg.message == "" {
		return errors.New("Message is null")
	}
	if len(msg.to) == 1 {
		msg.Pack = []byte(
			"To: " + msg.to[0] +
				"\r\n" + "Subject: " + msg.subject +
				"\r\n" + "From: Golosovanie <golosovanie.ru>" +
				"\r\n" + "Date: " + time.Now().Round(time.Second).Format(fmt.Sprintf(`"%s"`, time.RFC1123Z)) +
				// "\r\n" + "Content-Type: multipart/alternative; boundary=\"boundary-string\"" +
				// "\r\n" +
				"\r\n" + "Content-Type: text/plain; charset=\"utf-8\"" +
				"\r\n" + "Content-Transfer-Encoding: quoted-printable" +
				"\r\n" + "Content-Disposition: inline" +
				"\r\n" +
				"\r\n" + msg.message + "\r\n",
		)
		return nil
	}
	if len(msg.to) > 1 {
		for i := 0; i < len(msg.to); i++ {
			msg.Packs = append(msg.Packs, []byte(
				"To: "+msg.to[i]+
					"\r\n"+"Subject: "+msg.subject+
					"\r\n"+"From: Golosovanie <golosovanie.ru>"+
					"\r\n"+"Date: "+time.Now().Round(time.Second).Format(fmt.Sprintf(`"%s"`, time.RFC1123Z))+
					// "\r\n" + "Content-Type: multipart/alternative; boundary=\"boundary-string\"" +
					// "\r\n" +
					"\r\n"+"Content-Type: text/plain; charset=\"utf-8\""+
					"\r\n"+"Content-Transfer-Encoding: quoted-printable"+
					"\r\n"+"Content-Disposition: inline"+
					"\r\n"+
					"\r\n"+msg.message+"\r\n",
			),
			)
		}
		return nil
	}
	return nil
}

// func (msg *Message) method()  {

// }
