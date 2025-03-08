package proxy

import (
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"

	"github.com/or1ko/srpa/srpa/logging"
	"github.com/or1ko/srpa/srpa/session"
)

type ReverseProxyResource struct {
	Session *session.Session
	Logger  *logging.Logger
}

func (rp ReverseProxyResource) getSession(r *http.Request) (session.SessionInfo, bool) {
	return rp.Session.GetSession(r)
}

// Cookieを使用したリバースプロキシハンドラー
func (rp ReverseProxyResource) HandleReverseProxyWithCookieAuth(proxyPath string, target string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		// Cookieを取得して認証を検証
		sessionInfo, exists := rp.getSession(r)
		if !exists {
			http.Redirect(w, r, "/login?redirectTo="+r.URL.Path, http.StatusFound)
			// http.Error(w, "認証が必要です。ログインしてください。", http.StatusUnauthorized)
		}

		url, err := url.Parse(target)
		// リクエストを転送する先のURLを解析
		if err != nil {
			http.Error(w, "不正なターゲットURL", http.StatusInternalServerError)
			return
		}

		rp.Logger.Log(r.RemoteAddr, sessionInfo.Username, r.Method, r.RequestURI)

		proxy := httputil.NewSingleHostReverseProxy(url)
		r.URL.Path = strings.TrimPrefix(r.URL.Path, proxyPath)
		// r.Host = url.Host

		// リクエストの転送
		proxy.ServeHTTP(w, r)
	}
}
