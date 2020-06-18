package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/sorushsaghari/ie/internal/user"
	"net/http"
)

func IsAuthenticated() gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := c.Cookie("token")
		if err != nil {
			c.String(http.StatusUnauthorized, "token not provided")
			return
		}

		u, err := user.UserByToken(token)
		if err != nil {
			c.String(http.StatusUnauthorized, err.Error())
			return
		}

		c.Set("user", u)
		c.Next()
	}
}
