package usecase

import (
	"time"

	"github.com/mariofelesdossantosjunior/fc-microservices/internal/entity"
	"github.com/mariofelesdossantosjunior/fc-microservices/internal/gateway"
)

type CreateClientInputDTO struct {
	Name  string
	Email string
}

type CreateClientOutputDTO struct {
	ID        string
	Name      string
	Email     string
	CreateAt  time.Time
	UpdatedAt time.Time
}

type CreateClientUseCase struct {
	ClientGateway gateway.ClientGateway
}

func NewCreateClientUseCase(clientGateway gateway.ClientGateway) *CreateClientUseCase {
	return &CreateClientUseCase{
		ClientGateway: clientGateway,
	}
}

func (uc *CreateClientUseCase) Execute(input CreateClientInputDTO) (*CreateClientOutputDTO, error) {
	client, err := entity.NewClient(input.Name, input.Email)
	if err != nil {
		return nil, err
	}

	err = uc.ClientGateway.Save(client)
	if err != nil {
		return nil, err
	}

	output := &CreateClientOutputDTO{
		ID:        client.ID,
		Name:      client.Name,
		Email:     client.Email,
		CreateAt:  client.CreatedAt,
		UpdatedAt: client.UpdatedAt,
	}

	return output, nil
}
