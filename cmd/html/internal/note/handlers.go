package note

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/sorushsaghari/ie/internal/note"
	"github.com/sorushsaghari/ie/internal/user"
	"net/http"
	"strconv"
)

func Index(c echo.Context) error{
	u:= c.Get("user")
	notes, err := note.Find(u.(*user.User).ID)
	if err != nil {
		return c.Render(http.StatusBadRequest, "error.html", echo.Map{
			"Error": err.Error(),
		})

	}
	return c.Render(http.StatusOK, "list.html", echo.Map{
		"Title": "user list",
		"Notes": notes,
	})
}

func Detail(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	fmt.Println("test")
	if err != nil {
		return c.Render(http.StatusBadRequest, "error.html", echo.Map{
			"Error": err.Error(),
		})
	}
	n, err := note.One(uint(id))
	if err != nil {
		return c.Render(http.StatusBadRequest, "error.html", echo.Map{
			"Error": err.Error(),
		})

	}
	return c.Render(http.StatusOK, "detail.html", echo.Map{
		"Title": "user list",
		"Note": note.NewReadDto(*n),
	})
}

func Delete(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.Render(http.StatusBadRequest, "error.html", echo.Map{
			"Error": err.Error(),
		})
	}
	err = note.Delete(uint(id))
	if err != nil {
		return c.Render(http.StatusBadRequest, "error.html", echo.Map{
			"Error": err.Error(),
		})

	}

	return c.Redirect(http.StatusFound, "/note")
}

func Edit(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.Render(http.StatusBadRequest, "error.html", echo.Map{
			"Error": err.Error(),
		})
	}
	var body note.Dto
	c.Bind(body)
	err = note.Edit(uint(id), body.Parse())
	if err != nil {
		return c.Render(http.StatusBadRequest, "error.html", echo.Map{
			"Error": err.Error(),
		})

	}
	return c.Render(http.StatusOK, "detail.html", nil)
}