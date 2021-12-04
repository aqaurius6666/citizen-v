package lib

import (
	"net/http"

	"github.com/aqaurius6666/go-utils/utils"
	"github.com/gin-gonic/gin"
)

func BadRequest(c *gin.Context, err error) {
	c.Set("error", err)
	err = utils.Unwrap(err)
	c.AbortWithStatusJSON(http.StatusOK, gin.H{
		"success": false,
		"status":  400,
		"error":   err.Error(),
	})
}

func InternalServerError(c *gin.Context, err error) {
	c.Set("error", err)
	err = utils.Unwrap(err)
	c.AbortWithStatusJSON(http.StatusOK, gin.H{
		"success": false,
		"status":  500,
		"error":   err.Error(),
	})
}

func Success(c *gin.Context, response interface{}) {
	c.AbortWithStatusJSON(http.StatusOK, gin.H{
		"success": true,
		"status":  200,
		"data":    response,
	})
}

func Unauthorized(c *gin.Context, err error) {
	c.Set("error", err)
	err = utils.Unwrap(err)
	c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
		"success": false,
		"status":  401,
		"error":   err.Error(),
	})
}
