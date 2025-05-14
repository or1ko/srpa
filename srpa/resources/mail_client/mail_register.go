package mail_client

import (
	"bytes"
	_ "embed"
	"net/http"
	"text/template"
)

type MailRegisterResource struct {
	Host       string
	From       string
	Pool       MailPool
	MailClient Mail
	Home       string
}

func (res MailRegisterResource) MailRegisterHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {
		showMailRegisterPage(w)
		return
	}

	if r.Method != http.MethodPost {
		http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
		return
	}

	to_address := r.FormValue("mail")
	token, valid := res.Pool.AddSession(to_address)

	if valid {
		url := BuildMailPasswordUrl(res.Host, token)

		res.sendMail(to_address, url, res.From)

		showSendMatilPage(w, res.Home)

	} else {
		showInvalidMailPage(w)
	}
}

//go:embed mail.txt
var mail_text string

func (res MailRegisterResource) sendMail(address string, url string, from string) {
	var body bytes.Buffer

	t, _ := template.New("mail-text").Parse(mail_text)
	params := map[string]string{
		"From": from,
		"To":   address,
		"Url":  url,
	}
	t.Execute(&body, params)

	recivers := []string{address}

	res.MailClient.SendMail(recivers, body.Bytes())
}

//go:embed mail_register.html
var mail_register_page []byte

func showMailRegisterPage(w http.ResponseWriter) {
	w.Write(mail_register_page)
}

//go:embed invalid_mail_address.html
var invalid_mail_address_page []byte

func showInvalidMailPage(w http.ResponseWriter) {
	w.Write(invalid_mail_address_page)
}

//go:embed send_mail.html
var send_mail_page string

func showSendMatilPage(w http.ResponseWriter, home string) {
	t, _ := template.New("send_mail").Parse(send_mail_page)
	params := map[string]string{
		"Home": home,
	}
	t.Execute(w, params)
}

func RedirectMailRegisterPagee(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/mail_register", http.StatusFound)
}
