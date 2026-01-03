package createtransaction

import (
	"testing"

	"github.com/mariofelesdossantosjunior/fc-microservices/internal/entity"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type TransactionGatewayMock struct {
	mock.Mock
}

func (t *TransactionGatewayMock) Create(transaction *entity.Transaction) error {
	args := t.Called(transaction)
	return args.Error(0)
}

type AccountGatewayMock struct {
	mock.Mock
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

func TestCreateTransactionUseCase_Execute(t *testing.T) {
	client1, _ := entity.NewClient("Alice", "alice@example.com")
	account1 := entity.NewAccount(client1)
	account1.Credit(1000)

	client2, _ := entity.NewClient("Bob", "bob@example.com")
	account2 := entity.NewAccount(client2)
	account2.Credit(1000)

	accountGatewayMock := &MockAccountGateway{}
	accountGatewayMock.On("FindByID", account1.ID).Return(account1, nil)
	accountGatewayMock.On("FindByID", account2.ID).Return(account2, nil)

	transactionGatewayMock := &TransactionGatewayMock{}
	transactionGatewayMock.On("Create", mock.Anything).Return(nil)

	input := CreateTransactionInputDTO{
		AccountIDFrom: account1.ID,
		AccountIDTo:   account2.ID,
		Amount:        100,
	}

	useCase := NewCreateTransactionUseCase(transactionGatewayMock, accountGatewayMock)

	output, err := useCase.Execute(input)
	assert.Nil(t, err)
	assert.NotNil(t, output)
	assert.NotEmpty(t, output.ID)

	transactionGatewayMock.AssertExpectations(t)
	accountGatewayMock.AssertExpectations(t)
	accountGatewayMock.AssertNumberOfCalls(t, "FindByID", 2)
	transactionGatewayMock.AssertNumberOfCalls(t, "Create", 1)
}
