package logout

import (
	_ "embed"
	"net/http"

	"github.com/or1ko/srpa/srpa/resources/login"
	"github.com/or1ko/srpa/srpa/session"
)

type LogoutResource struct {
	Session *session.Session
	Home    string
}

func (res LogoutResource) LogoutHandler(w http.ResponseWriter, r *http.Request) {

	_, hasSession := res.Session.GetSession(r)
	if !hasSession {
		login.RedirectLoginPagee(w, r, "/logout")
		return
	}

	if r.Method == http.MethodGet {
		showChangePasswordPage(w)
		return
	}

	if r.Method != http.MethodPost {
		http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
		return
	}

	res.Session.RemoveSession(r)
	http.Redirect(w, r, res.Home, http.StatusFound)
}

//go:embed logout.html
var logout_page []byte

func showChangePasswordPage(w http.ResponseWriter) {
	w.Write(logout_page)
}
