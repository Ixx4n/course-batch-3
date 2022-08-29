package middleware

import (
	"course5-6/cmd/service"
	"strings"

	"github.com/gin-gonic/gin"
)

func WithAuthentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer") {
			c.JSON(401, map[string]string{
				"message": "unauthorize",
			})
			c.Abort()
		}

		auth := strings.Split(authHeader, " ")
		dataUser, err := service.DecryptJWT(auth[1])
		if err != nil {
			c.JSON(401, map[string]string{
				"message": "unauthorize",
			})
			c.Abort()
		}
		c.Set("id", dataUser["id"])
		c.Next()
	}
}
