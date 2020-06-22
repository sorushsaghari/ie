package note

import (
	"github.com/labstack/echo/v4"
	"github.com/sorushsaghari/ie/cmd/html/internal/middleware"
)

func Routers(r *echo.Echo) {
	note := r.Group("note")
	//note.Use()
	note.Use(middleware.IsAuthenticated)
	note.GET("", Index)
	note.GET("/:id", Detail)
	note.DELETE("/:id", Delete)
	//note.GET("/shit", func(c echo.Context) error{
	//	return c.Render(http.StatusOK, "list.html", map[string]interface{}{
	//		"Title": "list",
	//	})
	//})
	//note.GET("/:id",middleware.IsAuthenticated(), Detail)
}
