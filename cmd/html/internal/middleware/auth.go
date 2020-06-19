package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sorushsaghari/ie/internal/user"
	"net/http"
	"time"
)

func IsAuthenticated() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		fmt.Println("im here")
		token, err := c.Cookie("token")
		if err != nil {
			c.Redirect(http.StatusFound, "/user/login")
		}

		u, timeout, err := user.UserByToken(token)
		if err != nil {
			c.Redirect(http.StatusFound, "/user/login")
		}
		if time.Now().After(*timeout){
				c.Redirect(http.StatusFound, "/user/login")
		}
		c.Set("user", u)
	}
}
