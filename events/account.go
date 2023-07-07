
package events

import(
	queries "github.com/mohammadMghi/eventSourcing-SQRS/queries"
 
)
// Event Handler
type AccountEventHandler struct {
	QueryModel *queries.AccountQueryModel
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

type Event interface {
	// Marker interface for events
}

// Event Store
type EventStore struct {
	events []Event
}


// Event Bus
type EventBus struct {
	subscribers []AccountEventHandler
}


func (handler *AccountEventHandler) HandleEvent(event  Event) {
	switch event := event.(type) {
	case AccountCreatedEvent:
		// Update the query model for account creation event
		handler.QueryModel.CreateAccount(event.AccountNumber, event.AccountHolder, event.InitialBalance)
	case FundsDepositedEvent:
		// Update the query model for funds deposited event
		handler.QueryModel.DepositFunds(event.AccountNumber, event.Amount)
	}
}


func (store *EventStore) SaveEvent(event Event) {
	store.events = append(store.events, event)
}

func (bus *EventBus) Publish(event Event) {
	for _, subscriber := range bus.subscribers {
		subscriber.HandleEvent(event)
	}
}

func (bus *EventBus) Subscribe(subscriber AccountEventHandler) {
	bus.subscribers = append(bus.subscribers, subscriber)
}


