package account

import (
	"crypto/sha256"
	"fmt"
)

type Password struct {
	Password    string `json:"password"`
	IsEncrypted bool   `json:"encrypted"`
}

type Account struct {
	Id       string   `json:"id"`
	Role     string   `json:"role"`
	Password Password `json:"password"`
}

func (p Password) isMatch(password string) bool {
	if p.IsEncrypted {
		hashString := digest(password)
		return p.Password == hashString
	} else {
		return p.Password == password
	}
}

func digest(password string) string {
	hash := sha256.Sum256([]byte(password))
	hashString := fmt.Sprintf("%x", hash)
	return hashString
}

func (p Account) isMatchPassword(password string) bool {
	return p.Password.isMatch(password)
}

func ValueOf(s string) Password {
	return Password{Password: digest(s), IsEncrypted: true}
}

func (p *Account) ChangePassword(s string) {
	p.Password = ValueOf(s)
}
