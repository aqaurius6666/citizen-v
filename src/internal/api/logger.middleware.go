package api

import (
	"time"

	"github.com/aqaurius6666/boilerplate-server-go/src/internal/lib"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type LoggerMiddleware struct {
	logger *logrus.Logger
}

func (l *LoggerMiddleware) Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		_ = lib.SetBody(c)
		startTime := time.Now()
		c.Next()
		latencyTime := time.Since(startTime)
		reqMethod := c.Request.Method
		reqUri := c.Request.RequestURI
		statusCode := c.Writer.Status()
		reqLogger := l.logger.WithFields(logrus.Fields{
			"status":  statusCode,
			"latency": latencyTime,
			"method":  reqMethod,
			"path":    reqUri,
		})
		if body, ok := c.Get("body"); ok {
			reqLogger = reqLogger.WithField("body", string(body.([]byte)))
		}
		if err, ok := c.Get("error"); ok {
			reqLogger.Errorf("%+v", err)
		} else {
			reqLogger.Info()
		}
	}
}
