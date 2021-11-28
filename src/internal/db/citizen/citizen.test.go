package citizen

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/suite"
)

type Suite struct {
	suite.Suite
	Repository CitizenRepo
	Mock       sqlmock.Sqlmock
}

func (s *Suite) SelectCitizen() {

}

func (s *Suite) InsertCitizen() {

}
