package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewClient(t *testing.T) {
	name := "John Doe"
	email := "john@doe.com"

	client, err := NewClient(name, email)
	assert.Nil(t, err)
	assert.NotNil(t, client)
	assert.Equal(t, name, client.Name)
	assert.Equal(t, email, client.Email)
}

func TestCreateNewClientWhenAtgsAreInvalid(t *testing.T) {
	name := ""
	email := "invalid-email"

	client, err := NewClient(name, email)
	assert.Nil(t, client)
	assert.NotNil(t, err)
}

func TestUpdateClient(t *testing.T) {
	name := "John Doe"
	email := "john@doe.com"
	client, _ := NewClient(name, email)

	newName := "Jane Doe"
	newEmail := "jane@doe.com"
	err := client.Update(newName, newEmail)
	assert.Nil(t, err)
	assert.Equal(t, newName, client.Name)
	assert.Equal(t, newEmail, client.Email)
}

func TestUpdateClientWhenAtgsAreInvalid(t *testing.T) {
	name := "John Doe"
	email := "john@doe.com"
	client, _ := NewClient(name, email)

	newName := ""
	newEmail := "invalid-email"
	err := client.Update(newName, newEmail)
	assert.NotNil(t, err)
	assert.Equal(t, name, client.Name)
	assert.Equal(t, email, client.Email)
}

func TestAddAccountToClient(t *testing.T) {
	name := "John Doe"
	email := "john@doe.com"
	client, _ := NewClient(name, email)
	account := NewAccount(client)
	err := client.AddAccount(account)
	assert.Nil(t, err)
	assert.Equal(t, account, client.Accounts[0])
}

func TestAddAccountToClientWhenAccountDoesNotBelongToClient(t *testing.T) {
	name1 := "John Doe"
	email1 := "john@doe.com"
	client1, _ := NewClient(name1, email1)

	name2 := "Jane Doe"
	email2 := "jane@doe.com"
	client2, _ := NewClient(name2, email2)
	account := NewAccount(client2)
	err := client1.AddAccount(account)
	assert.NotNil(t, err)
	assert.Equal(t, 0, len(client1.Accounts))
}
