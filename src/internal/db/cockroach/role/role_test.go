package role

import (
	"context"
	"database/sql"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/aquarius6666/citizen-v/src/internal/db/role"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DBSuite struct {
	role.RoleSuite
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
	repo, err := InitRoleCDBMockRepo(context.Background(), s.DB)
	s.RoleSuite = role.RoleSuite{
		Repository: repo,
		Mock:       mock,
	}
	assert.Nil(s.T(), err)
}
func (s *DBSuite) TestSelectRole() {
	s.RoleSuite.SelectRole()

}

func TestSelectUser(t *testing.T) {
	suite.Run(t, new(DBSuite))
}
