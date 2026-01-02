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
