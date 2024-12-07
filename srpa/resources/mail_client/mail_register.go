package mail_client

import (
	"bytes"
	_ "embed"
	"net/http"
	"text/template"
)

type MailRegisterResource struct {
	Host       string
	Pool       MailPool
	MailClient Mail
}

func (res MailRegisterResource) MailRegisterHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {
		showMailRegisterPage(w, r)
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

		res.sendMail(to_address, url)

		showSendMatilPage(w, r)

	} else {
		showInvalidMailPage(w, r)
	}
}

//go:embed mail.txt
var mail_text string

func (res MailRegisterResource) sendMail(address string, url string) {
	var body bytes.Buffer

	t, _ := template.New("mail-text").Parse(mail_text)
	params := map[string]string{
		"To":  address,
		"Url": url,
	}
	t.Execute(&body, params)

	recivers := []string{address}

	res.MailClient.SendMail(recivers, body.Bytes())
}

//go:embed mail_register.html
var mail_register_page []byte

func showMailRegisterPage(w http.ResponseWriter, r *http.Request) {
	w.Write(mail_register_page)
}

//go:embed invalid_mail_address.html
var invalid_mail_address_page []byte

func showInvalidMailPage(w http.ResponseWriter, r *http.Request) {
	w.Write(invalid_mail_address_page)
}

//go:embed send_mail.html
var send_mail_page []byte

func showSendMatilPage(w http.ResponseWriter, r *http.Request) {
	w.Write(send_mail_page)
}

func RedirectMailRegisterPagee(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/mail_register", http.StatusFound)
}
