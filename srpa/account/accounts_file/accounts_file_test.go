package accounts_file

import (
	"testing"

	"github.com/or1ko/srpa/srpa/account"
)

func TestSave(t *testing.T) {
	p := account.Account{
		Id:       "test",
		Password: account.ValueOf("pass"),
	}
	k := map[string]account.Account{
		p.Id: p,
	}
	a := account.Accounts{
		Accounts: k,
	}

	file := AccountsFile{
		Filename: "users.json",
		Accounts: a,
	}

	file.Save()

}
