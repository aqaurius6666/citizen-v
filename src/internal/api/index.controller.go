package api

import (
	"math/rand"

	"github.com/aquarius6666/citizen-v/src/internal/db"
	"github.com/aquarius6666/citizen-v/src/internal/db/user"
	"github.com/aquarius6666/citizen-v/src/internal/lib"
	"github.com/gin-gonic/gin"
)

type IndexController struct {
	Repo db.ServerRepo
}

func (i *IndexController) HandleIndexGet(g *gin.Context) {

	user, _ := i.Repo.SelectUser(&user.Search{})
	_, err := i.Repo.InsertUser(user)
	if err != nil {
		lib.BadRequest(g, err)
		return
	}
	lib.Success(g, "Go go bruh bruh ...")
}

func (i *IndexController) HandleRandomGet(g *gin.Context) {
	rand := rand.Int()
	g.JSON(200, gin.H{"random": rand})
}
