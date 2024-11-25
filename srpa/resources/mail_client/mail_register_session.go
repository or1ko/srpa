package mail_client

import (
	"crypto/rand"
	"encoding/base64"
	"log"
)

type MailPool struct {
	validSession       map[string]string
	validMailAddresses []string
}

func ValueOf(addresses []string) MailPool {
	return MailPool{
		validMailAddresses: addresses,
		validSession:       make(map[string]string),
	}
}

func (s MailPool) AddSession(mail string) (string, bool) {
	isValid := validMail(s, mail)
	if isValid {
		token := generateToken()
		s.validSession[token] = mail

		return token, true
	} else {
		return "", false
	}

}

func validMail(s MailPool, mail string) bool {
	if mail == "" {
		return false
	}

	isValid := false
	for temp := range s.validMailAddresses {
		if s.validMailAddresses[temp] == mail {
			isValid = true
		}
	}
	return isValid
}

func (s MailPool) Valid(token string) (string, bool) {
	mail, exists := s.validSession[token]
	return mail, exists
}

func generateToken() string {
	b := make([]byte, 32)
	_, err := rand.Read(b)
	if err != nil {
		log.Fatal("Failed to generate token")
	}
	return base64.URLEncoding.EncodeToString(b)
}
