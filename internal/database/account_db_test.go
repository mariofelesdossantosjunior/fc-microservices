package database

import (
	"database/sql"
	"testing"

	"github.com/mariofelesdossantosjunior/fc-microservices/internal/entity"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/suite"
)

type AccountDBTestSuite struct {
	suite.Suite
	db        *sql.DB
	accountDB *AccountDB
	client    *entity.Client
}

func (suite *AccountDBTestSuite) SetupTest() {
	var err error
	suite.db, err = sql.Open("sqlite3", ":memory:")
	suite.Require().NoError(err)

	_, err = suite.db.Exec(`
	CREATE TABLE clients (
		id TEXT PRIMARY KEY,
		name TEXT,
		email TEXT,
		created_at DATETIME
	);
	`)

	_, err = suite.db.Exec(`
	CREATE TABLE accounts (
		id TEXT PRIMARY KEY,
		client_id TEXT,
		balance REAL,
		created_at DATETIME,
		FOREIGN KEY (client_id) REFERENCES clients(id)
	);
	`)

	suite.accountDB = NewAccountDB(suite.db)
	suite.client, _ = entity.NewClient("John Doe", "j@j.com")

}

func (suite *AccountDBTestSuite) TearDownTest() {
	suite.db.Close()
	suite.db.Exec("DROP TABLE accounts")
	suite.db.Exec("DROP TABLE clients")
}

func TestAccountDBTestSuite(t *testing.T) {
	suite.Run(t, new(AccountDBTestSuite))
}

func (suite *AccountDBTestSuite) TestSave() {
	account := entity.NewAccount(suite.client)

	err := suite.accountDB.Save(account)
	suite.Require().NoError(err)
}

func (suite *AccountDBTestSuite) TestFindByID() {
	suite.db.Exec("INSERT INTO clients (id, name, email, created_at) VALUES (?, ?, ?, ?)",
		suite.client.ID, suite.client.Name, suite.client.Email, suite.client.CreatedAt)

	account := entity.NewAccount(suite.client)
	err := suite.accountDB.Save(account)
	suite.Nil(err)

	retrievedAccount, err := suite.accountDB.Get(account.ID)
	suite.Require().NoError(err)
	suite.Equal(account.ID, retrievedAccount.ID)
	suite.Equal(account.Client.ID, retrievedAccount.Client.ID)
	suite.Equal(account.Balance, retrievedAccount.Balance)
}
