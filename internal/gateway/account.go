package gateway

import "github.com/mariofelesdossantosjunior/fc-microservices/internal/entity"

type AccountGateway interface {
	FindByID(id string) (*entity.Account, error)
	Save(account *entity.Account) error
}
