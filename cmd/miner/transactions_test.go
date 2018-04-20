package main

import (
	"strings"
	"testing"
)

func TestNewTransactions(t *testing.T) {
	transactions := NewTransactions(NewEvents())
	if len(transactions.pool) != 0 {
		t.Errorf("NewTransactions(): Failed, couldn't find pool.")
	}
}

func TestTransactionsCreate(t *testing.T) {
	payload := `{"payload": "Arnulf Beckenbauer"}`
	transactions := NewTransactions(NewEvents())
	transaction, _ := transactions.Create(strings.NewReader(payload))

	if transaction.Payload != "Arnulf Beckenbauer" {
		t.Errorf("Transactions#Create(): Payload wasn't defined.")
	}
}

func TestTransactionsAdd(t *testing.T) {
	payload := `{"payload": "Arnulf Beckenbauer"}`
	transactions := NewTransactions(NewEvents())
	transaction, _ := transactions.Create(strings.NewReader(payload))
	transactions.Add(*transaction)

	if transactions.pool[0].Id != transaction.Id {
		t.Errorf("Transactions#Add(): Couldn't find transaction in transactions pool")
	}
}

func TestTransactionsGet(t *testing.T) {
	payload := `{"payload": "Arnulf Beckenbauer"}`
	transactions := NewTransactions(NewEvents())
	transaction, _ := transactions.Create(strings.NewReader(payload))
	transactions.Add(*transaction)

	actual := transactions.Get(transaction.Id)

	if actual.Id != transaction.Id {
		t.Errorf("Transactions#Add(): Couldn't find transaction in transactions pool")
	}
}

func TestTransactionsPop(t *testing.T) {
	events := NewEvents()
	transactions := NewTransactions(events)
	ts := transactions.Pop()

	if len(ts) != 0 {
		t.Errorf("Transactions#pop(): Number of transactions should be 0.")
	}
}

func TestTransactionsConfirm(t *testing.T) {
	events := NewEvents()
	events.Start()
	transactions := NewTransactions(events)

	for i := 0; i < 10; i++ {
		payload := `{"payload": "Arnulf Beckenbauer"}`
		transaction, _ := transactions.Create(strings.NewReader(payload))
		transactions.Add(*transaction)
	}

	id := transactions.pool[0].Id
	transactions.Confirm([]Transaction{transactions.pool[0]})

	if transactions.pool[0].Id == id {
		t.Errorf("Transactions#Confirm(): Transaction was found in transactions pool, but should've been confirmed.")
	}
}
