package api

import (
	"github.com/aqaurius6666/citizen-v/src/internal/lib"
	"github.com/aqaurius6666/citizen-v/src/internal/var/e"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type RoleMiddleware struct {
	logger *logrus.Logger
}

func (l *RoleMiddleware) OnlyRole(role string) gin.HandlerFunc {
	return func(c *gin.Context) {
		ro, ok := c.Get("role")
		if !ok || ro != role {
			lib.Unauthorized(c, e.ErrAuthNoPermission)
			return
		}
		c.Next()
	}
}
