package database

import (
	"database/sql"
	"testing"

	"github.com/mariofelesdossantosjunior/fc-microservices/internal/entity"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/suite"
)

type ClientDBTestSuite struct {
	suite.Suite
	db       *sql.DB
	clientDB *ClientDB
}

func (suite *ClientDBTestSuite) SetupTest() {
	var err error
	suite.db, err = sql.Open("sqlite3", ":memory:")
	suite.Require().NoError(err)

	suite.clientDB = NewClientDB(suite.db)

	_, err = suite.db.Exec(`
	CREATE TABLE clients (
		id TEXT PRIMARY KEY,
		name TEXT,
		email TEXT,
		created_at DATETIME
	);
	`)
	suite.Require().NoError(err)
}

func (suite *ClientDBTestSuite) TearDownTest() {
	suite.db.Close()
	suite.db.Exec("DROP TABLE clients")
}

func TestClientDBTestSuite(t *testing.T) {
	suite.Run(t, new(ClientDBTestSuite))
}

func (suite *ClientDBTestSuite) TestSaveAndGetClient() {
	client, _ := entity.NewClient("John Doe", "john@example.com")

	err := suite.clientDB.Save(client)
	suite.Require().NoError(err)

	retrievedClient, err := suite.clientDB.Get(client.ID)
	suite.Require().NoError(err)
	suite.Equal(client.ID, retrievedClient.ID)
	suite.Equal(client.Name, retrievedClient.Name)
	suite.Equal(client.Email, retrievedClient.Email)
}
