package note

import (
	"github.com/gin-gonic/gin"
	"github.com/sorushsaghari/ie/internal/note"
	"github.com/sorushsaghari/ie/internal/user"
	"net/http"
	"strconv"
)

func Index(c *gin.Context){
	u, _ := c.Get("user")

	notes, err := note.Find(u.(user.User).ID)
	if err != nil {
		c.HTML(http.StatusBadRequest, "error.html", gin.H{
			"Error": err.Error(),
		})
		return
	}
	c.HTML(http.StatusOK, "list.html", gin.H{
		"Title": "user list",
		"Notes": notes,
	})
	return
}

func Detail(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.HTML(http.StatusBadRequest, "error.html", gin.H{
			"Error": err.Error(),
		})
		return
	}
	note, err := note.One(uint(id))
	if err != nil {
		c.HTML(http.StatusBadRequest, "error.html", gin.H{
			"Error": err.Error(),
		})
		return
	}
	c.HTML(http.StatusOK, "detail.html", gin.H{
		"Title": "user list",
		"Notes": note,
	})
	return
}