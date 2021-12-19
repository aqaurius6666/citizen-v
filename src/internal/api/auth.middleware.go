package api

import (
	"strings"

	"github.com/aqaurius6666/citizen-v/src/internal/lib"
	"github.com/aqaurius6666/citizen-v/src/internal/services/jwt"
	"github.com/aqaurius6666/citizen-v/src/internal/var/e"
	"github.com/gin-gonic/gin"
)

type AuthMiddleware struct {
	JWTService jwt.JWT
}

func (s *AuthMiddleware) CheckAuth(g *gin.Context) {
	var authString string
	if authString = g.GetHeader("Authorization"); authString == "" {
		lib.Unauthorized(g, e.ErrAuthMissingAuthorization)
		return
	}
	if !strings.HasPrefix(authString, "Bearer ") {
		lib.Unauthorized(g, e.ErrAuthTokenInvalid)
		return
	}
	token := authString[7:]
	ok, data, err := s.JWTService.Verify(token)
	if err != nil {
		lib.Unauthorized(g, err)
		return
	}
	if !ok {
		lib.Unauthorized(g, e.ErrAuthTokenFail)
		return
	}
	g.Set("uid", data["uid"])
	g.Set("role", data["role"])
	g.Next()
}
