package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateTransaction(t *testing.T) {
	client1, _ := NewClient("Alice", "alice@example.com")
	account1 := NewAccount(client1)
	client2, _ := NewClient("Bob", "bob@example.com")
	account2 := NewAccount(client2)

	account1.Credit(200.0)
	account2.Credit(50.0)

	transaction, err := NewTransaction(account1, account2, 75.0)
	assert.Nil(t, err)
	assert.NotNil(t, transaction)
	assert.Equal(t, 125.0, account1.Balance)
	assert.Equal(t, 125.0, account2.Balance)
}

func TestCreateTransactionWithInsufficientFunds(t *testing.T) {
	client1, _ := NewClient("Alice", "alice@example.com")
	account1 := NewAccount(client1)
	client2, _ := NewClient("Bob", "bob@example.com")
	account2 := NewAccount(client2)

	account1.Credit(200.0)
	account2.Credit(50.0)

	transaction, err := NewTransaction(account1, account2, 300.0)
	assert.NotNil(t, err)
	assert.Nil(t, transaction)
	assert.Equal(t, 200.0, account1.Balance)
	assert.Equal(t, 50.0, account2.Balance)
}
