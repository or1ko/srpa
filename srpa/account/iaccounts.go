package account

type IAccounts interface {
	Add(string, string) bool
	ChangePassword(string, string)
	Confirm(string, string) (Account, bool)
}
