package database

import (
	"database/sql"
	"testing"

	"github.com/mariofelesdossantosjunior/fc-microservices/internal/entity"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/suite"
)

type TransactionDBTestSuite struct {
	suite.Suite
	db            *sql.DB
	client1       *entity.Client
	client2       *entity.Client
	accountFrom   *entity.Account
	accountTo     *entity.Account
	transactionDB *TransactionDB
}

func (suite *TransactionDBTestSuite) SetupTest() {
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

	_, err = suite.db.Exec(`
	CREATE TABLE transactions (
		id TEXT PRIMARY KEY,
		account_id_from TEXT,
		account_id_to TEXT,
		amount REAL,
		created_at DATETIME,
		FOREIGN KEY (account_id_from) REFERENCES accounts(id),
		FOREIGN KEY (account_id_to) REFERENCES accounts(id)
	);
	`)

	client1, err := entity.NewClient("John Doe", "john@doe.com")
	suite.client1 = client1
	suite.Require().NoError(err)

	client2, err := entity.NewClient("Jane Doe", "jane@doe.com")
	suite.client2 = client2
	suite.Require().NoError(err)

	accountFrom := entity.NewAccount(suite.client1)
	accountFrom.Balance = 1000.0
	suite.accountFrom = accountFrom
	suite.Require().NoError(err)

	accountTo := entity.NewAccount(suite.client2)
	accountTo.Balance = 500.0
	suite.accountTo = accountTo
	suite.Require().NoError(err)

	suite.transactionDB = NewTransactionDB(suite.db)

}

func (suite *TransactionDBTestSuite) TearDownTest() {
	suite.db.Close()
	suite.db.Exec("DROP TABLE transactions")
	suite.db.Exec("DROP TABLE accounts")
	suite.db.Exec("DROP TABLE clients")
}

func TestTransactionDBTestSuite(t *testing.T) {
	suite.Run(t, new(TransactionDBTestSuite))
}

func (suite *TransactionDBTestSuite) TestCreateTransaction() {
	transaction, err := entity.NewTransaction(suite.accountFrom, suite.accountTo, 200.0)
	suite.Require().NoError(err)

	err = suite.transactionDB.Create(transaction)
	suite.Require().NoError(err)
}
