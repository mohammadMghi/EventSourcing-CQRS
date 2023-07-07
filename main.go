package main

import "fmt"

// Command Model
type CreateAccountCommand struct {
	AccountNumber string
	AccountHolder string
	InitialBalance float64
}

type DepositFundsCommand struct {
	AccountNumber string
	Amount        float64
}

// Command Handler
type AccountCommandHandler struct {
	eventStore EventStore
	eventBus   EventBus
}
type Event interface {
	// Marker interface for events
}

func (handler *AccountCommandHandler) CreateAccount(command CreateAccountCommand) {
	// Apply validations and business rules

	// Generate event
	event := AccountCreatedEvent{
		AccountNumber:  command.AccountNumber,
		AccountHolder:  command.AccountHolder,
		InitialBalance: command.InitialBalance,
	}

	// Persist event to the event store
	handler.eventStore.SaveEvent(event)

	// Publish event to the event bus
	handler.eventBus.Publish(event)
}

func (handler *AccountCommandHandler) DepositFunds(command DepositFundsCommand) {
	// Apply validations and business rules

	// Generate event
	event := FundsDepositedEvent{
		AccountNumber: command.AccountNumber,
		Amount:        command.Amount,
	}

	// Persist event to the event store
	handler.eventStore.SaveEvent(event)

	// Publish event to the event bus
	handler.eventBus.Publish(event)
}

// Event Store
type EventStore struct {
	events []Event
}

func (store *EventStore) SaveEvent(event Event) {
	store.events = append(store.events, event)
}

// Event Bus
type EventBus struct {
	subscribers []AccountEventHandler
}

func (bus *EventBus) Publish(event Event) {
	for _, subscriber := range bus.subscribers {
		subscriber.HandleEvent(event)
	}
}

func (bus *EventBus) Subscribe(subscriber AccountEventHandler) {
	bus.subscribers = append(bus.subscribers, subscriber)
}

// Event Handler
type AccountEventHandler struct {
	queryModel *AccountQueryModel
}

func (handler *AccountEventHandler) HandleEvent(event Event) {
	switch event := event.(type) {
	case AccountCreatedEvent:
		// Update the query model for account creation event
		handler.queryModel.CreateAccount(event.AccountNumber, event.AccountHolder, event.InitialBalance)
	case FundsDepositedEvent:
		// Update the query model for funds deposited event
		handler.queryModel.DepositFunds(event.AccountNumber, event.Amount)
	}
}

// AccountCreatedEvent
type AccountCreatedEvent struct {
	AccountNumber  string
	AccountHolder  string
	InitialBalance float64
}

// FundsDepositedEvent
type FundsDepositedEvent struct {
	AccountNumber string
	Amount        float64
}


// Query Model
type AccountQueryModel struct {
	accounts map[string]Account
}

type Account struct {
	AccountNumber string
	AccountHolder string
	Balance       float64
}

func (model *AccountQueryModel) CreateAccount(accountNumber, accountHolder string, initialBalance float64) {
	account := Account{
		AccountNumber: accountNumber,
		AccountHolder: accountHolder,
		Balance:       initialBalance,
	}
	model.accounts[accountNumber] = account
}

func (model *AccountQueryModel) DepositFunds(accountNumber string, amount float64) {
	account := model.accounts[accountNumber]
	account.Balance += amount
	model.accounts[accountNumber] = account
}

// Example usage
func main() {
	eventStore := EventStore{}
	eventBus := EventBus{}
	queryModel := AccountQueryModel{
		accounts: make(map[string]Account),
	}
	eventHandler := AccountEventHandler{queryModel: &queryModel}

	commandHandler := AccountCommandHandler{
		eventStore: eventStore,
		eventBus:   eventBus,
	}
	eventBus.Subscribe(eventHandler)

	// Create account command
	createCommand := CreateAccountCommand{
		AccountNumber: "123456789",
		AccountHolder: "John Doe",
		InitialBalance: 1000,
	}
	commandHandler.CreateAccount(createCommand)

	// Deposit funds command
	depositCommand := DepositFundsCommand{
		AccountNumber: "123456789",
		Amount:        500,
	}
	commandHandler.DepositFunds(depositCommand)

	// Retrieve account balance
	account := queryModel.accounts["123456789"]
	fmt.Println("Account Balance:", account.Balance) // Output: 1500
}