package api

import (
	"github.com/aqaurius6666/boilerplate-server-go/src/internal/db"
	"github.com/gin-gonic/gin"
)

type IndexController struct {
	Repo db.ServerRepo
}

func (i *IndexController) HandleIndexGet(g *gin.Context) {
	g.JSON(200, gin.H{"message": "Go go bruh bruh ..."})
}
