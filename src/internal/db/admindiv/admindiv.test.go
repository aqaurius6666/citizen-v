package admindiv

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/suite"
)

type Suite struct {
	suite.Suite
	Repository AdminDivRepo
	Mock       sqlmock.Sqlmock
}

func (s *Suite) SelectAdminDiv() {

}

func (s *Suite) InsertAdminDiv() {

}
