package admindiv

import (
	"context"
	"database/sql"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/aquarius6666/citizen-v/src/internal/db/admindiv"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DBSuite struct {
	admindiv.Suite
	DB *gorm.DB
}

func (s *DBSuite) SetupSuite() {
	var (
		db   *sql.DB
		err  error
		mock sqlmock.Sqlmock
	)

	db, mock, err = sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	assert.Nil(s.T(), err)
	s.DB, err = gorm.Open(postgres.New(
		postgres.Config{Conn: db},
	), &gorm.Config{
		SkipDefaultTransaction: true,
		PrepareStmt:            false,
		DisableAutomaticPing:   true,
	})
	assert.Nil(s.T(), err)
	repo, err := InitAdminDivCDBMockRepo(context.Background(), s.DB)
	s.Suite = admindiv.Suite{
		Repository: repo,
		Mock:       mock,
	}
	assert.Nil(s.T(), err)
}
func (s *DBSuite) TestAdminDiv() {
	s.Suite.SelectAdminDiv()

}

func TestRun(t *testing.T) {
	suite.Run(t, new(DBSuite))
}
