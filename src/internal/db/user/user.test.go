package user

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type UserSuite struct {
	suite.Suite
	Repository UserRepo
	Mock       sqlmock.Sqlmock
}

func (s *UserSuite) SelectUser() {
	// u1 := User{}
	// usr, err := s.Repository.InsertUser(&u1)
	// assert.Nil(s.T(), err)
	s.Mock.ExpectQuery("SELECT * FROM \"users\" ORDER BY \"users\".\"id\" LIMIT 1")
	new_usr, err := s.Repository.SelectUser(&Search{})
	assert.Nil(s.T(), err)
	s.T().Log(new_usr)
	// assert.Equal(s.T(), usr.ID, new_usr.ID)
	// s.T().Log(usr.ID, new_usr.ID)
}

func (s *UserSuite) InsertUser() {
	usr, err := s.Repository.InsertUser(&User{})
	assert.Nil(s.T(), err)
	new_usr, err := s.Repository.SelectUser(&Search{})
	assert.Nil(s.T(), err)
	assert.Equal(s.T(), usr.ID, new_usr.ID)
	s.T().Log(usr.ID, new_usr.ID)
}
