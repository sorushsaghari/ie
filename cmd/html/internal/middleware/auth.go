package middleware

import (
	"github.com/labstack/echo/v4"
	"github.com/sorushsaghari/ie/internal/user"
	"net/http"
	"time"
)
func IsAuthenticated(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		token, err := c.Cookie("token")
		if err != nil {
			return c.Redirect(http.StatusFound, "/user/login")
		}

		u, timeout, err := user.UserByToken(token.Value)
		if err != nil {
			return c.Redirect(http.StatusFound, "/user/login")

		}
		if time.Now().After(*timeout){
			return c.Redirect(http.StatusFound, "/user/login")
		}
		c.Set("user", u)
		return next(c)
	}
}
