package api

import (
	"math/rand"

	"github.com/aqaurius6666/citizen-v/src/internal/db"
	"github.com/aqaurius6666/citizen-v/src/internal/lib"
	"github.com/gin-gonic/gin"
)

type IndexController struct {
	Repo db.ServerRepo
}

func (i *IndexController) HandleIndexGet(g *gin.Context) {
	lib.Success(g, "Go go bruh bruh ...")
}

func (i *IndexController) HandleRandomGet(g *gin.Context) {
	lib.Success(g, rand.Int())
}
