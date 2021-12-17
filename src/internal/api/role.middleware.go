package api

import (
	"github.com/aqaurius6666/citizen-v/src/internal/lib"
	"github.com/aqaurius6666/citizen-v/src/internal/model"
	"github.com/aqaurius6666/citizen-v/src/internal/var/e"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

type RoleMiddleware struct {
	logger *logrus.Logger
	Model  model.Server
}

func (l *RoleMiddleware) OnlyRole(role ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		ro, ok := c.Get("role")
		if !ok {
			lib.Unauthorized(c, e.ErrAuthNoPermission)
			return
		}
		valid := false
		for _, r := range role {
			if r == ro.(string) {
				valid = true
			}
		}
		if !valid {
			lib.Unauthorized(c, e.ErrAuthNoPermission)
			return
		}
		c.Next()
	}
}

func (l *RoleMiddleware) OnlyActive() gin.HandlerFunc {
	return func(c *gin.Context) {
		uid, err := uuid.Parse(c.GetString("uid"))
		if err != nil {
			lib.Unauthorized(c, e.ErrIdInvalid)
			return
		}

		ok, err := l.Model.IsRoleActive(uid)
		if err != nil || !ok {
			lib.Unauthorized(c, e.ErrAuthNoPermission)
			return
		}
		c.Next()
	}
}
