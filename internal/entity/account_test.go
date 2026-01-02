package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateNewAccount(t *testing.T) {
	name := "John Doe"
	email := "john.doe@example.com"
	client, _ := NewClient(name, email)
	account := NewAccount(client)
	if account == nil {
		t.Error("Expected account to be created")
	}
	if account.Client != client {
		t.Error("Expected account client to be the same as the client passed in")
	}
	if account.Balance != 0.0 {
		t.Error("Expected account balance to be 0.0")
	}
	assert.NotNil(t, account)
	assert.Equal(t, client.ID, account.Client.ID)
}

func TestCreateNewAccountWhenClientIsNil(t *testing.T) {
	account := NewAccount(nil)
	assert.Nil(t, account)
}

func TestCreditAccount(t *testing.T) {
	name := "John Doe"
	email := "john.doe@example.com"
	client, _ := NewClient(name, email)
	account := NewAccount(client)
	account.Credit(100.0)
	assert.Equal(t, float64(100.0), account.Balance)
}

func TestDebitAccount(t *testing.T) {
	name := "John Doe"
	email := "john.doe@example.com"
	client, _ := NewClient(name, email)
	account := NewAccount(client)
	account.Credit(100.0)
	account.Debit(40.0)
	assert.Equal(t, float64(60.0), account.Balance)
}
