package user

import "C"
import (
	"github.com/labstack/echo/v4"
	"github.com/sorushsaghari/ie/internal/user"
	"net/http"
	"time"
)

func Register(c echo.Context) error {
	var body user.CreateDto
	if err := c.Bind(&body); err != nil {
		return c.Render(http.StatusBadRequest, "error.Render", map[string]interface{}{
			"Error": err.Error(),
		})
		//return
	}
	if err := user.Create(body.Parse()); err != nil {
		return c.Render(http.StatusBadRequest, "error.Render", map[string]interface{}{
			"Error": err.Error(),
		})

	}
	return c.String(http.StatusOK, "done")

}

func RegisterPage(c echo.Context) error{
	return c.Render(http.StatusOK, "user.Render", map[string]interface{}{})
}

func LoginPage(c echo.Context) error{
	return c.Render(http.StatusOK, "login.html", map[string]interface{}{
		"TItle": "login",
	})
}
func Login(c echo.Context) error{
	var body user.Dto
	if err := c.Bind(&body); err != nil {
		return c.Render(http.StatusBadRequest, "error.Render", map[string]interface{}{
			"Error": err.Error(),
		})
	}
	u, err := user.One(body.Parse())
	if err != nil || u == nil {
		return c.Render(http.StatusBadRequest, "error.Render", map[string]interface{}{
			"Error": err.Error(),
		})
	}
	tomorrow := time.Now().AddDate(0, 0, 1)
	auth := user.Auth{
		Token: user.TokenGenerator(),
		UserID: u.ID,
		TimeOut: &tomorrow,
	}
	if err = user.Store(&auth);err != nil {
		return c.Render(http.StatusBadRequest, "error.Render", map[string]interface{}{
			"Error": err.Error(),
		})

	}

	c.SetCookie(&http.Cookie{
		Name:       "token",
		Value:      auth.Token,
		Path:       "/",
		Domain:     "127.0.0.1",
		Expires:    time.Now().AddDate(0, 0, 1),
		Secure:     false,
		HttpOnly:   true,
	})
	return c.String(http.StatusOK, "log in was successful")

}
