package middlewares

import (
	"net/http"

	"github.com/Bobby-P-dev/FinalProject2_kel7/helpers"
	"github.com/gin-gonic/gin"
)

func Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		verifToken, err := helpers.VerifToken(c)
		_ = verifToken

		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "Unautenticated",
				"message": err.Error(),
			})
			return
		}
		c.Set("userData", verifToken)
		c.Next()
	}
}

func UserAuthorization() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}
