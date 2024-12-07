package add_user

import (
	_ "embed"
	"net/http"

	"github.com/or1ko/srpa/srpa/account"
	"github.com/or1ko/srpa/srpa/resources/login"
	"github.com/or1ko/srpa/srpa/session"
)

type AddUserResource struct {
	Accounts account.IAccounts
	Session  *session.Session
}

func (res AddUserResource) AddUserHandler(w http.ResponseWriter, r *http.Request) {

	sessionInfo, hasSession := res.Session.GetSession(r)
	if !hasSession {
		login.RedirectLoginPagee(w, r, "/add_user")
		return
	}

	if r.Method == http.MethodGet {
		ShowAddUserPage(w)
		return
	}

	if r.Method != http.MethodPost {
		http.Error(w, "Invalid Method", http.StatusMethodNotAllowed)
		return
	}

	if sessionInfo.Role != "admin" {
		ShowNoPermissionPage(w)
		return
	}

	user := r.FormValue("username")
	pass := r.FormValue("password")

	res.Accounts.Add(user, pass)
	http.Redirect(w, r, "/add_user", http.StatusFound)
}

//go:embed add_user.html
var add_user_html []byte

func ShowAddUserPage(w http.ResponseWriter) {
	w.Write(add_user_html)
}

//go:embed no_permission.html
var no_permission_page []byte

func ShowNoPermissionPage(w http.ResponseWriter) {
	w.Write(no_permission_page)
}
