package user

import (
	"github.com/labstack/echo/v4"
)

func Routers(r *echo.Echo) {
	user := r.Group("/user")
	user.GET("/", RegisterPage)
	user.POST("/", Register)
	user.GET("/login", LoginPage)
	user.POST("/login", Login)
	//user.Use(middleware.IsAuthenticated())
}
