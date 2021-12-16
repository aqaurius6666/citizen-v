package model

import (
	"context"
	"testing"

	"github.com/aqaurius6666/citizen-v/src/internal/db"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestHasPermission(t *testing.T) {

	ctx := context.Background()
	logger := logrus.New()
	repo, err := db.InitServerRepo(ctx, logger, db.DBDsn("postgresql://root:root@cdb:26257/defaultdb?sslmode=disable"))
	assert.Nil(t, err)

	model, err := NewServerModel(ctx, logger, repo)
	assert.Nil(t, err)

	testcase := []map[string]interface{}{
		{
			"u": uuid.MustParse("8993850e-9445-4aec-946c-fdb63698739d"),
			"a": uuid.MustParse("6a31b531-1bb5-496a-b00f-d7cf6161bf6f"),
			"e": true,
		},
	}
	for _, s := range testcase {
		res, err := model.HasPermission(s["u"].(uuid.UUID), s["a"].(uuid.UUID))
		assert.Nil(t, err)
		assert.Equal(t, s["e"], res, "case (%s, %s)", s["u"], s["a"])
	}
}
