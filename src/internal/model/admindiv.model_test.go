package model

import (
	"testing"

	ct "github.com/aqaurius6666/citizen-v/src/internal/var/t"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestGetNewCode(t *testing.T) {

	model := SetupModel()
	if !assert.NotNil(t, model) {
		return
	}

	testcase := []map[string]interface{}{
		{
			"a": ct.A011001, //a011001
			"e": "00000000",
		},
		{
			"a": ct.A0110, //a0110
			"e": "000000",
		},
		{
			"a": ct.A011001, //a011001
			"e": "00000000",
		},
		{
			"a": uuid.Nil, //00000
			"e": "00",
		},
		{
			"a": ct.A03, //a03
			"e": "0000",
		},
	}

	for _, s := range testcase {
		res, err := model.GetNewCode(s["a"].(uuid.UUID))
		if !assert.Nil(t, err) {
			return
		}
		if !assert.Equal(t, len(s["e"].(string)), len(res), "case (%s)", s["a"]) {
			return
		}
	}
}

func TestIsChild(t *testing.T) {
	model := SetupModel()
	if !assert.NotNil(t, model) {
		return
	}
	testcase := []map[string]interface{}{
		{
			"a": "",
			"c": []string{"01"},
			"e": true,
		},
	}

	for _, s := range testcase {
		res := model.IsChild(s["a"].(string), s["c"].([]string))
		if !assert.Equal(t, s["e"], res, "case (%s)", s["a"]) {
			return
		}
	}

}
