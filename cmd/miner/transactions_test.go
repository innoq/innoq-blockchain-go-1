package main

import "testing"

func TestNewTransactions(t *testing.T) {
	transactions := NewTransactions()
	if len(transactions.pool) != 0 {
		t.Errorf("NewTransactions(): Failed, couldn't find pool.")
	}
}

func TestTransactionsCreate(t *testing.T) {
	payload := `{"payload": "Arnulf Beckenbauer"}`
	transactions := NewTransactions()
	transaction := transactions.Create(payload)

	if transaction.Confirmed {
		t.Errorf("Transactions#Create(): Expected Confirmed to be false, was true.")
	}

	if transaction.Payload != payload {
		t.Errorf("Transactions#Create(): Payload wasn't defined.")
	}
}

func TestTransactionsAdd(t *testing.T) {
	payload := `{"payload": "Arnulf Beckenbauer"}`
	transactions := NewTransactions()
	transaction := transactions.Create(payload)
	transactions.Add(*transaction)

	if transactions.pool[0].Id != transaction.Id {
		t.Errorf("Transactions#Add(): Couldn't find transaction in transactions pool")
	}
}

func TestTransactionsGet(t *testing.T) {
	payload := `{"payload": "Arnulf Beckenbauer"}`
	transactions := NewTransactions()
	transaction := transactions.Create(payload)
	transactions.Add(*transaction)

	actual := transactions.Get(transaction.Id)

	if actual.Id != transaction.Id {
		t.Errorf("Transactions#Add(): Couldn't find transaction in transactions pool")
	}
}
