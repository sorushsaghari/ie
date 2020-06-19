package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/sorushsaghari/ie/internal/user"
	"net/http"
	"time"
)

func IsAuthenticated() gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := c.Cookie("token")
		if err != nil {
			c.Redirect(http.StatusUnauthorized, "user/login")
			return
		}

		u, timeout, err := user.UserByToken(token)
		if err != nil {
			c.Redirect(http.StatusUnauthorized, "user/login")
			return
		}
		if time.Now().After(*timeout){
				c.Redirect(http.StatusUnauthorized, "user/login")
				return
		}
		c.Set("user", u)
		c.Next()
	}
}
