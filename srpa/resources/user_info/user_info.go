package user_info

import (
	_ "embed"
	"html/template"
	"net/http"

	"github.com/or1ko/srpa/srpa/resources/login"
	"github.com/or1ko/srpa/srpa/session"
)

type UserInfoResource struct {
	Session *session.Session
	Home    string
}

func (res UserInfoResource) UserInfoHandler(w http.ResponseWriter, r *http.Request) {

	sessionInfo, hasSession := res.Session.GetSession(r)
	if !hasSession {
		login.RedirectLoginPagee(w, r, "/user_info")
		return
	}

	showUserInfoPage(w, sessionInfo, res.Home)
}

//go:embed user_info.html
var user_info_page string

func showUserInfoPage(w http.ResponseWriter, info session.SessionInfo, home string) {
	w.Header().Set("Content-Type", "text/html")
	t, _ := template.New("user-info").Parse(user_info_page)
	params := map[string]string{
		"Username":      info.Username,
		"Role":          info.Role,
		"LastAcceeTime": info.LastAccessTime.Format("2006-01-02T15:04:05Z07:00"),
		"Home":          home,
	}
	t.Execute(w, params)
}
