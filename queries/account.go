package queries

import(
	modles "github.com/mohammadMghi/eventSourcing-CQRS/models"
)
// Query Model
type AccountQueryModel struct {
	Accounts map[string]modles.Account
}

func (model *AccountQueryModel) CreateAccount(accountNumber, accountHolder string, initialBalance float64) {
	account := modles.Account{
		AccountNumber: accountNumber,
		AccountHolder: accountHolder,
		Balance:       initialBalance,
	}
	model.Accounts[accountNumber] = account
}

func (model *AccountQueryModel) DepositFunds(accountNumber string, amount float64) {
	account := model.Accounts[accountNumber]
	account.Balance += amount
	model.Accounts[accountNumber] = account
}
