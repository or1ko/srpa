package add_user

import (
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
		ShowAddUserPage(w, r)
		return
	}

	if r.Method != http.MethodPost {
		http.Error(w, "Invalid Method", http.StatusMethodNotAllowed)
		return
	}

	if sessionInfo.Role != "admin" {
		ShowNoPermissionPage(w, r)
		return
	}

	user := r.FormValue("username")
	pass := r.FormValue("password")

	res.Accounts.Add(user, pass)
	http.Redirect(w, r, "/add_user", http.StatusFound)
}

func ShowAddUserPage(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "frontend/add_user.html")
}

func ShowNoPermissionPage(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "frontend/no_permission.html")
}
