package account

type Accounts struct {
	Accounts map[string]Account `json:"accounts"`
}

func (a Accounts) ToAccountsValues() []Account {
	return values(a.Accounts)
}

func values(m map[string]Account) []Account {

	var values []Account

	for _, value := range m {
		values = append(values, value)
	}

	return values

}

func (as *Accounts) Add(username string, password string) bool {
	account := Account{
		Id:       username,
		Password: ValueOf(password),
		Role:     "general",
	}
	return as.add(account)
}

func (a *Accounts) ChangePassword(username string, password string) {
	a.Accounts[username] = Account{Id: username, Password: ValueOf(password)}
}

func (a Accounts) Confirm(username string, password string) (Account, bool) {
	account := a.Accounts[username]
	success := account.isMatchPassword(password)
	if success {
		return account, true
	} else {
		return Account{}, false
	}
}

func (as *Accounts) add(a Account) bool {
	_, exists := as.Accounts[a.Id]
	if !exists {
		as.Accounts[a.Id] = a
	}
	return !exists
}

func (as *Accounts) Remove(username string) bool {
	_, exists := as.Accounts[username]
	if exists {
		delete(as.Accounts, username)
	}
	return exists
}

func (as *Accounts) Get(username string) (Account, bool) {
	account, exists := as.Accounts[username]
	if exists {
		return account, true
	} else {
		return Account{}, false
	}
}
