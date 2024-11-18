package emailx

import (
	"crypto/tls"
	"fmt"
	"github.com/jordan-wright/email"
	"net/smtp"
)

type EmailClient struct {
	conf Conf
}

func NewClient(conf Conf) *EmailClient {
	return &EmailClient{conf}
}

func (x *EmailClient) Send(to []string, subject string, body string) error {
	auth := smtp.PlainAuth("", x.conf.From, x.conf.Secret, x.conf.Host)
	e := email.NewEmail()
	if x.conf.Nickname != "" {
		e.From = fmt.Sprintf("%s <%s>", x.conf.Nickname, x.conf.From)
	} else {
		e.From = x.conf.From
	}
	e.To = to
	e.Subject = subject
	e.HTML = []byte(body)
	var err error
	hostAddr := fmt.Sprintf("%s:%d", x.conf.Host, x.conf.Port)
	if x.conf.IsSSL {
		err = e.SendWithTLS(hostAddr, auth, &tls.Config{ServerName: x.conf.Host})
	} else {
		err = e.Send(hostAddr, auth)
	}

	return err
}
