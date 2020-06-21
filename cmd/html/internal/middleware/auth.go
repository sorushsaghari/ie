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
			c.HTML(http.StatusUnauthorized, "login.html", gin.H{
				"Title": "login",
			})
			return
		}

		u, timeout, err := user.UserByToken(token)
		if err != nil {
			c.HTML(http.StatusUnauthorized, "login.html", gin.H{
				"Title": "login",
			})
			return
		}
		if time.Now().After(*timeout){
			c.HTML(http.StatusUnauthorized, "login.html", gin.H{
				"Title": "login",
			})
			return
		}
		c.Set("user", u)
		return
	}
}
