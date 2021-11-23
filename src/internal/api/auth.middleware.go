package api

import (
	"github.com/aquarius6666/citizen-v/src/internal/lib"
	"github.com/aquarius6666/citizen-v/src/internal/services/jwt"
	"github.com/aquarius6666/citizen-v/src/internal/var/e"
	"github.com/gin-gonic/gin"
)

type AuthMiddleware struct {
	JWTService jwt.JWTService
}

func (s *AuthMiddleware) CheckAuth(g *gin.Context) {
	ok, err := s.JWTService.Verify()
	if err != nil {
		lib.Unauthorized(g, err)
		return
	}
	if !ok {
		lib.Unauthorized(g, e.ErrAuthTokenFail)
	}
	g.Next()
}
