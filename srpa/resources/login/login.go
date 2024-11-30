package login

import (
	_ "embed"
	"html/template"
	"net/http"

	"github.com/or1ko/srpa/srpa/account"
	"github.com/or1ko/srpa/srpa/session"
)

type LoginResource struct {
	Accounts account.IAccounts
	Session  *session.Session
}

func (res LoginResource) LoginHandler(w http.ResponseWriter, r *http.Request) {

	rediretUrl := r.URL.Query().Get("redirectTo")
	if r.Method == http.MethodGet {
		showLoginPage(w, rediretUrl)
		return
	}

	if r.Method != http.MethodPost {
		http.Error(w, "Invalid Method", http.StatusMethodNotAllowed)
		return
	}

	user := r.FormValue("username")
	pass := r.FormValue("password")

	account, success := res.Accounts.Confirm(user, pass)
	if success {
		res.Session.AddSession(w, r, account)
		http.Redirect(w, r, rediretUrl, http.StatusFound)
		return
	}
	ShowLoginFailurePage(w, r)
}

//go:embed login.html
var login_html string

func showLoginPage(w http.ResponseWriter, redirectUrl string) {
	w.Header().Set("Content-Type", "text/html")
	t, _ := template.New("loginpage").Parse(login_html)
	t.Execute(w, redirectUrl)
}

func RedirectLoginPagee(w http.ResponseWriter, r *http.Request, redirectUrl string) {
	http.Redirect(w, r, "/login?redirectTo="+redirectUrl, http.StatusFound)
}

//go:embed login_failure.html
var login_failure_page []byte

func ShowLoginFailurePage(w http.ResponseWriter, r *http.Request) {
	w.Write(login_failure_page)
}
