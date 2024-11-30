package mail_client

import (
	_ "embed"
	"net/http"
	"time"

	"github.com/or1ko/srpa/srpa/account"
)

type MailPasswordResource struct {
	ExpiredMinute int
	Pool          MailPool
	CookieName    string
	Accounts      account.IAccounts
}

func (res MailPasswordResource) MailPasswordHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {
		token := r.URL.Query().Get("token")

		_, valid := res.Pool.Valid(token)

		if valid {
			http.SetCookie(w, &http.Cookie{
				Name:     res.CookieName,
				Value:    token,
				Path:     "/mail_password",
				Expires:  time.Now().Add(time.Duration(res.ExpiredMinute) * time.Minute),
				HttpOnly: true,
			})
			showMailPasswordPage(w, r)
			return
		}

		RedirectMailRegisterPagee(w, r)
		return
	}

	if r.Method != http.MethodPost {
		http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
		return
	}

	cookie, err := r.Cookie(res.CookieName)

	if err != nil {
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	token := cookie.Value
	mail, valid := res.Pool.Valid(token)
	if valid {
		pass := r.FormValue("password")
		res.Accounts.Add(mail, pass)
		showSuccessMailPasswordPage(w, r)
		return
	}
}

func BuildMailPasswordUrl(host string, token string) string {
	return host + "/mail_password?token=" + token
}

//go:embed mail_password.html
var mail_password_page []byte

func showMailPasswordPage(w http.ResponseWriter, r *http.Request) {
	w.Write(mail_password_page)
}

//go:embed success_mail_password.html
var success_mail_password_page []byte

func showSuccessMailPasswordPage(w http.ResponseWriter, r *http.Request) {
	w.Write(success_mail_password_page)
}
