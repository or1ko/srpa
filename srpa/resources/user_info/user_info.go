package user_info

import (
	"html/template"
	"net/http"

	"github.com/or1ko/srpa/srpa/resources/login"
	"github.com/or1ko/srpa/srpa/session"
)

type UserInfoResource struct {
	Session *session.Session
}

func (res UserInfoResource) UserInfoHandler(w http.ResponseWriter, r *http.Request) {

	sessionInfo, hasSession := res.Session.GetSession(r)
	if !hasSession {
		login.RedirectLoginPagee(w, r, "/user_info")
		return
	}

	showUserInfoPage(w, sessionInfo)
}

func showUserInfoPage(w http.ResponseWriter, info session.SessionInfo) {
	w.Header().Set("Content-Type", "text/html")
	t, _ := template.ParseFiles("frontend/user_info.html")
	params := map[string]string{
		"Username":      info.Username,
		"Role":          info.Role,
		"LastAcceeTime": info.LastAccessTime.Format("2006-01-02T15:04:05Z07:00"),
	}
	t.Execute(w, params)
}
