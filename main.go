package main

import (
	events "github.com/mohammadMghi/eventSourcing-SQRS/events"
	commands "github.com/mohammadMghi/eventSourcing-SQRS/commands"
	modles "github.com/mohammadMghi/eventSourcing-SQRS/models"
	queries "github.com/mohammadMghi/eventSourcing-SQRS/queries"
	"fmt"
 
)



// Example usage
func main() {
	eventStore := events.EventStore{}
	eventBus := events.EventBus{}
	queryModel := queries.AccountQueryModel{
		Accounts: make(map[string]modles.Account),
	}
	eventHandler := events.AccountEventHandler{QueryModel: &queryModel}

	commandHandler := commands.AccountCommandHandler{
		EventStore: eventStore,
		EventBus:   eventBus,
	}
	eventBus.Subscribe(eventHandler)

	// Create account command
	createCommand := commands.CreateAccountCommand{
		AccountNumber: "123456789",
		AccountHolder: "John Doe",
		InitialBalance: 1000,
	}
	commandHandler.CreateAccount(createCommand)

	// Deposit funds command
	depositCommand := commands.DepositFundsCommand{
		AccountNumber: "123456789",
		Amount:        500,
	}
	commandHandler.DepositFunds(depositCommand)

	// Retrieve account balance
	account := queryModel.Accounts["123456789"]
	fmt.Println("Account Balance:", account.Balance) // Output: 1500
}