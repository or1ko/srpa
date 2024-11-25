package proxy

import (
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/or1ko/srpa/srpa/session"
)

type ReverseProxyResource struct {
	Session *session.Session
}

func (rp ReverseProxyResource) hasSession(r *http.Request) bool {
	return rp.Session.HasSession(r)
}

// Cookieを使用したリバースプロキシハンドラー
func (rp ReverseProxyResource) HandleReverseProxyWithCookieAuth(target string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		// Cookieを取得して認証を検証
		hasSession := rp.hasSession(r)
		if !hasSession {
			http.Redirect(w, r, "/login?redirectTo="+r.URL.Path, http.StatusFound)
			// http.Error(w, "認証が必要です。ログインしてください。", http.StatusUnauthorized)
		}

		url, err := url.Parse(target)
		// リクエストを転送する先のURLを解析
		if err != nil {
			http.Error(w, "不正なターゲットURL", http.StatusInternalServerError)
			return
		}

		// リバースプロキシを作成
		proxy := httputil.NewSingleHostReverseProxy(url)
		r.Host = url.Host

		// リクエストの転送
		proxy.ServeHTTP(w, r)
	}
}
