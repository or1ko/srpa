package mail_client

import (
	"bytes"
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

	mail := r.FormValue("mail")
	token, valid := res.Pool.AddSession(mail)

	if valid {
		url := BuildMailPasswordUrl(res.Host, token)

		var body bytes.Buffer

		t, _ := template.ParseFiles("mail/mail.txt")
		params := map[string]string{
			"To":  mail,
			"Url": url,
		}
		t.Execute(&body, params)

		recivers := []string{mail}

		res.MailClient.SendMail(recivers, body.Bytes())
		// メール送る
		// showTestUrl(w, url)
		showSendMatilPage(w, r)

	} else {
		showInvalidMailPage(w, r)
	}
}

func showMailRegisterPage(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "frontend/mail_register.html")
}

func showInvalidMailPage(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "frontend/invalid_mail_address.html")
}

func showSendMatilPage(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "frontend/send_mail.html")
}

func RedirectMailRegisterPagee(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/mail_register", http.StatusFound)
}
