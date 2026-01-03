package createaccount

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

type MockAccountGateway struct {
	mock.Mock
}

func (m *MockAccountGateway) FindByID(id string) (*entity.Account, error) {
	args := m.Called(id)
	return args.Get(0).(*entity.Account), args.Error(1)
}

func (m *MockAccountGateway) Save(account *entity.Account) error {
	args := m.Called(account)
	return args.Error(0)
}

func TestCreateAccountUseCase_Execute(t *testing.T) {
	client, _ := entity.NewClient("John Doe", "john@example.com")

	mClient := &MockClientGateway{}
	mClient.On("Find", client.ID).Return(client, nil)

	mAccount := &MockAccountGateway{}
	mAccount.On("Save", mock.Anything).Return(nil)

	uc := NewCreateAccountUseCase(mAccount, mClient)

	input := CreateAccountInputDTO{
		ClientID: client.ID,
	}

	output, err := uc.Execute(input)
	assert.Nil(t, err)
	assert.NotNil(t, output)
	assert.NotEmpty(t, output.ID)
	mClient.AssertExpectations(t)
	mClient.AssertNumberOfCalls(t, "Find", 1)
	mAccount.AssertExpectations(t)
	mAccount.AssertNumberOfCalls(t, "Save", 1)
}
