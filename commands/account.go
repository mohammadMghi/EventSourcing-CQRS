package commands


import(
	events "github.com/mohammadMghi/eventSourcing-SQRS/events"
 
)
// Command Model
type CreateAccountCommand struct {
	AccountNumber string
	AccountHolder string
	InitialBalance float64
}
// Command Handler
type AccountCommandHandler struct {
	EventStore events.EventStore
	EventBus   events.EventBus
}

type DepositFundsCommand struct {
	AccountNumber string
	Amount        float64
}


func (handler *AccountCommandHandler) CreateAccount(command CreateAccountCommand) {
	// Apply validations and business rules

	// Generate event
	event := events.AccountCreatedEvent{
		AccountNumber:  command.AccountNumber,
		AccountHolder:  command.AccountHolder,
		InitialBalance: command.InitialBalance,
	}

	// Persist event to the event store
	handler.EventStore.SaveEvent(event)

	// Publish event to the event bus
	handler.EventBus.Publish(event)
}

func (handler *AccountCommandHandler) DepositFunds(command DepositFundsCommand) {
	// Apply validations and business rules

	// Generate event
	event := events.FundsDepositedEvent{
		AccountNumber: command.AccountNumber,
		Amount:        command.Amount,
	}

	// Persist event to the event store
	handler.EventStore.SaveEvent(event)

	// Publish event to the event bus
	handler.EventBus.Publish(event)
}