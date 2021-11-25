package role

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/suite"
)

type RoleSuite struct {
	suite.Suite
	Repository RoleRepo
	Mock       sqlmock.Sqlmock
}

func (s *RoleSuite) SelectRole() {

}

func (s *RoleSuite) InsertRole() {

}
