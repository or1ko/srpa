package mail_client

import "net/smtp"

type Mail struct {
	Addr string
	From string
	Auth smtp.Auth
}

func (m Mail) SendMail(to []string, msg []byte) error {
	error := smtp.SendMail(m.Addr, m.Auth, m.From, to, msg)
	return error
}
