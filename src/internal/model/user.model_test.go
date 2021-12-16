package model

import (
	"context"
	"testing"

	"github.com/aqaurius6666/citizen-v/src/internal/db"
	ct "github.com/aqaurius6666/citizen-v/src/internal/var/t"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func SetupModel() Server {
	ctx := context.Background()
	logger := logrus.New()
	repo, err := db.InitServerRepo(ctx, logger, db.DBDsn("postgresql://root:root@localhost:20000/defaultdb?sslmode=disable"))
	if err != nil {
		return nil
	}

	model, err := NewServerModel(ctx, logger, repo)
	if err != nil {
		return nil
	}
	return model
}

func TestHasPermission(t *testing.T) {

	model := SetupModel()
	if !assert.NotNil(t, model) {
		return
	}
	testcase := []map[string]interface{}{
		{
			"u": ct.U011001, //u011001
			"a": ct.A011001, //a011001
			"e": true,
		},
		{
			"u": ct.U011001, //u011001
			"a": ct.A0110,   //a0110
			"e": false,
		},
		{
			"u": ct.U01,     //u01
			"a": ct.A011001, //a011001
			"e": true,
		},
		{
			"u": ct.U01,   //u01
			"a": uuid.Nil, //00000
			"e": false,
		},
		{
			"u": ct.U0110, //u0110
			"a": ct.A03,   //a03
			"e": false,
		},
	}
	for _, s := range testcase {
		res, err := model.HasPermission(s["u"].(uuid.UUID), s["a"].(uuid.UUID))
		if !assert.Nil(t, err) {
			return
		}
		if !assert.Equal(t, s["e"], res, "case (%s, %s)", s["u"], s["a"]) {
			return
		}
	}
}

func TestIsRoleActive(t *testing.T) {
	model := SetupModel()
	if !assert.NotNil(t, model) {
		return
	}
	testcase := []map[string]interface{}{
		{
			"u": ct.U01, //u011001
			"e": true,
		},
		{
			"u": uuid.Nil,
			"e": false,
		},
		{
			"u": ct.U03,
			"e": false,
		},
		{
			"u": ct.U0301,
			"e": false,
		},
	}
	for _, s := range testcase {
		res, err := model.IsRoleActive(s["u"].(uuid.UUID))
		if !assert.Nil(t, err) {
			return
		}
		if !assert.Equal(t, s["e"], res, "case (%s)", s["u"]) {
			return
		}
	}
}
