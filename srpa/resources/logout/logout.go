package logout

import (
	"net/http"

	"github.com/or1ko/srpa/srpa/resources/login"
	"github.com/or1ko/srpa/srpa/session"
)

type LogoutResource struct {
	Session *session.Session
}

func (res LogoutResource) LogoutHandler(w http.ResponseWriter, r *http.Request) {

	_, hasSession := res.Session.GetSession(r)
	if !hasSession {
		login.RedirectLoginPagee(w, r, "/logout")
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

	res.Session.RemoveSession(r)
	http.Redirect(w, r, "/", http.StatusFound)
}

func showChangePasswordPage(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "frontend/logout.html")
}
