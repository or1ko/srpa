package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"strings"

	"github.com/or1ko/srpa/srpa/account/accounts_file"
	"github.com/or1ko/srpa/srpa/config"
	"github.com/or1ko/srpa/srpa/proxy"
	"github.com/or1ko/srpa/srpa/resources/add_user"
	"github.com/or1ko/srpa/srpa/resources/login"
	"github.com/or1ko/srpa/srpa/resources/logout"
	"github.com/or1ko/srpa/srpa/resources/mail_client"
	"github.com/or1ko/srpa/srpa/resources/password"
	"github.com/or1ko/srpa/srpa/resources/user_info"
	"github.com/or1ko/srpa/srpa/session"
)

func main() {
	workdir := flag.String("workdir", "", "spacify work directory")

	flag.Parse()

	if *workdir != "" {
		os.Chdir(*workdir)
	}

	wd, err := os.Getwd()
	if err == nil {
		fmt.Println("Current workdir is " + wd)
	}

	config := config.Load("config.yaml")

	ipAddress, err := getLocalIP()
	if err != nil {
		log.Fatalf("Failed to get local IP address: %v\n", err)
	}
	host := "http://" + ipAddress + ":" + config.Port

	as := accounts_file.Load("users.json")
	session := session.EmptySession()
	logger := config.Logging.NewLogger()
	login := login.LoginResource{
		Accounts: as,
		Session:  &session,
	}

	password := password.PasswordResource{
		Session:  session,
		Accounts: &as,
	}

	proxy := proxy.ReverseProxyResource{
		Session: &session,
		Logger:  &logger,
	}

	logout := logout.LogoutResource{
		Session: &session,
	}

	add_user := add_user.AddUserResource{
		Accounts: as,
		Session:  &session,
	}

	user_info := user_info.UserInfoResource{
		Session: &session,
	}

	mail_pool := mail_client.ValueOf(config.Mail.MailAddress)

	mail_registerr := mail_client.MailRegisterResource{
		Host: host,
		Pool: mail_pool,
		MailClient: mail_client.Mail{
			Addr: config.Mail.Host,
			From: config.Mail.From,
			Auth: nil,
		},
	}

	mail_password := mail_client.MailPasswordResource{
		Pool:          mail_pool,
		Accounts:      as,
		CookieName:    "mail_cookie",
		ExpiredMinute: 10,
	}

	http.HandleFunc("/login", login.LoginHandler)
	http.HandleFunc("/password", password.ChangePasswordHandler)
	http.HandleFunc("/logout", logout.LogoutHandler)
	http.HandleFunc("/add_user", add_user.AddUserHandler)
	http.HandleFunc("/user_info", user_info.UserInfoHandler)
	http.HandleFunc("/mail_register", mail_registerr.MailRegisterHandler)
	http.HandleFunc("/mail_password", mail_password.MailPasswordHandler)

	for i := 0; i < len(config.ReverseMaps); i++ {
		paths := strings.SplitN(config.ReverseMaps[i], ":", 2)
		log.Printf("%s %s", paths[0], paths[1])
		http.HandleFunc(paths[0], proxy.HandleReverseProxyWithCookieAuth(paths[0], paths[1]))
	}

	log.Printf("Starting reverse proxy server on %s\n", host)

	log.Fatal(http.ListenAndServe(":"+config.Port, nil))
}

func getLocalIP() (string, error) {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		return "", err
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)
	return localAddr.IP.String(), nil
}
