package password

import (
	_ "embed"
	"net/http"

	"github.com/or1ko/srpa/srpa/account"
	"github.com/or1ko/srpa/srpa/resources/login"
	"github.com/or1ko/srpa/srpa/session"
)

type PasswordResource struct {
	Accounts account.IAccounts
	Session  session.Session
}

func (res PasswordResource) ChangePasswordHandler(w http.ResponseWriter, r *http.Request) {

	sessionInfo, hasSession := res.Session.GetSession(r)
	if !hasSession {
		login.RedirectLoginPagee(w, r, "/password")
		return
	}

	if r.Method == http.MethodGet {
		showChangePasswordPage(w, r)
		return
	}

	if r.Method != http.MethodPost {
		http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
		return
	}

	pass := r.FormValue("password")
	res.Accounts.ChangePassword(sessionInfo.Username, pass)
	http.Redirect(w, r, "/", http.StatusFound)

}

//go:embed password.html
var password_page []byte

func showChangePasswordPage(w http.ResponseWriter, r *http.Request) {
	w.Write(password_page)
}
