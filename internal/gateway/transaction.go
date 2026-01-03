package gateway

import "github.com/mariofelesdossantosjunior/fc-microservices/internal/entity"

type TransactionGateway interface {
	Create(transaction *entity.Transaction) error
}
