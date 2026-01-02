package gateway

import "github.com/mariofelesdossantosjunior/fc-microservices/internal/entity"

type ClientGateway interface {
	Find(id string) (*entity.Client, error)
	Save(client *entity.Client) error
}
