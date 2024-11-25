package accounts_file

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/or1ko/srpa/srpa/account"
)

type AccountsFile struct {
	Filename string
	Accounts account.Accounts
}

func Load(filename string) AccountsFile {
	bytes, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	var accounts []account.Account
	if err := json.Unmarshal(bytes, &accounts); err != nil {
		log.Fatal(err)
	}

	file := AccountsFile{
		Filename: filename,
		Accounts: account.Accounts{Accounts: toMap(accounts)},
	}
	return file
}

func toMap(accounts []account.Account) map[string]account.Account {
	m := map[string]account.Account{}

	for _, a := range accounts {
		m[a.Id] = a
	}

	return m
}

func (a AccountsFile) Save() {
	prettyJSON, err := json.MarshalIndent(a.Accounts.ToAccountsValues(), "", "  ")
	if err != nil {
		fmt.Println("Error marshaling to JSON:", err)
		return
	}

	file, err := os.Create(a.Filename)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	_, err = file.Write(prettyJSON)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}
}

func (as AccountsFile) Add(username string, password string) bool {
	success := as.Accounts.Add(username, password)
	as.Save()
	return success
}

func (as AccountsFile) ChangePassword(username string, password string) {
	as.Accounts.ChangePassword(username, password)
	as.Save()
}

func (as AccountsFile) Confirm(username string, password string) (account.Account, bool) {
	return as.Accounts.Confirm(username, password)
}
