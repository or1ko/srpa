package session

import (
	"crypto/rand"
	"encoding/base64"
	"log"
	"net/http"
	"time"

	"github.com/or1ko/srpa/srpa/account"
)

type Session struct {
	SessionLifeTimeMinute int
	cookieName            string
	validSession          map[string]SessionInfo
}

func EmptySession() Session {
	return Session{
		SessionLifeTimeMinute: 30,
		cookieName:            "session_token",
		validSession:          make(map[string]SessionInfo),
	}
}

func (s Session) HasSession(r *http.Request) bool {
	_, exists := s.GetSession(r)
	return exists
}

func (s Session) GetSession(r *http.Request) (SessionInfo, bool) {
	cookie, err := r.Cookie(s.cookieName)
	if err != nil {
		return SessionInfo{}, false
	} else {
		key := cookie.Value
		sessionInfo, exists := s.validSession[key]

		if exists {
			validTime := isValidTime(sessionInfo, s.SessionLifeTimeMinute)
			if validTime {
				sessionInfo.UpdateLastAccessTime()
				return sessionInfo, true
			}
			return SessionInfo{}, false
		} else {
			return SessionInfo{}, false
		}
	}
}

func isValidTime(sessionInfo SessionInfo, lifeTimeMinute int) bool {
	sessionExpiredTime := sessionInfo.ExpiredTime(lifeTimeMinute)
	return time.Now().Before(sessionExpiredTime)
}

func (s Session) AddSession(w http.ResponseWriter, r *http.Request, account account.Account) {
	token := generateToken()
	s.validSession[token] = ValueOf(account)
	http.SetCookie(w, &http.Cookie{
		Name:     s.cookieName,
		Value:    token,
		Path:     "/",
		Expires:  time.Now().Add(24 * time.Hour),
		HttpOnly: true,
	})
}

func (s Session) RemoveSession(r *http.Request) bool {
	cookie, err := r.Cookie(s.cookieName)
	if err != nil {
		return true
	} else {
		key := cookie.Value
		delete(s.validSession, key)
		return true
	}
}

func generateToken() string {
	b := make([]byte, 32)
	_, err := rand.Read(b)
	if err != nil {
		log.Fatal("Failed to generate token")
	}
	return base64.URLEncoding.EncodeToString(b)
}
