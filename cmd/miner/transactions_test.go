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
