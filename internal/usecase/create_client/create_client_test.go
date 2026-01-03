package usecase

import (
	"testing"

	"github.com/mariofelesdossantosjunior/fc-microservices/internal/entity"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockClientGateway struct {
	mock.Mock
}

func (m *MockClientGateway) Find(id string) (*entity.Client, error) {
	args := m.Called(id)
	return args.Get(0).(*entity.Client), args.Error(1)
}

func (m *MockClientGateway) Save(client *entity.Client) error {
	args := m.Called(client)
	return args.Error(0)
}

func TestCreateClientUseCase_Execute(t *testing.T) {
	m := &MockClientGateway{}
	m.On("Save", mock.Anything).Return(nil)
	uc := NewCreateClientUseCase(m)

	input := CreateClientInputDTO{
		Name:  "John Doe",
		Email: "john@doe.com",
	}

	output, err := uc.Execute(input)
	assert.Nil(t, err)
	assert.NotNil(t, output)
	assert.NotEmpty(t, output.ID)
	assert.Equal(t, input.Name, output.Name)
	m.AssertExpectations(t)
	m.AssertNumberOfCalls(t, "Save", 1)
}
