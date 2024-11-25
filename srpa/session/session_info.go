package session

import (
	"time"

	"github.com/or1ko/srpa/srpa/account"
)

type SessionInfo struct {
	Username string
	Role     string

	LastAccessTime time.Time
}

func ValueOf(account account.Account) SessionInfo {
	return SessionInfo{
		Username:       account.Id,
		Role:           account.Role,
		LastAccessTime: time.Now(),
	}
}

func (si *SessionInfo) UpdateLastAccessTime() {
	si.LastAccessTime = time.Now()
}

func (si *SessionInfo) ExpiredTime(SessionLifeTimeMinute int) time.Time {
	return si.LastAccessTime.Add(time.Duration(SessionLifeTimeMinute) * time.Minute)
}
