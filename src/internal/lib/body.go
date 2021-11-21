package lib

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/aqaurius6666/boilerplate-server-go/src/internal/var/e"
	"github.com/gin-gonic/gin"
)

func SetBody(c *gin.Context) error {
	if c.Request.Method == http.MethodGet {
		return nil
	}
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		return e.ErrMissingBody
	}
	c.Set("body", body)
	return nil
}

func GetBody(g *gin.Context, target interface{}) error {
	var bbody []byte
	var err error
	ibody, ok := g.Get("body")
	if !ok {
		return e.ErrMissingBody
	} else {
		bbody = ibody.([]byte)
	}
	err = json.Unmarshal(bbody, target)
	if err != nil {
		return nil
	}
	return nil
}
