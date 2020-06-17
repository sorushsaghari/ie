package user

import (
	"github.com/gin-gonic/gin"
	"github.com/sorushsaghari/ie/internal/user"
	"net/http"
)

func Register(c *gin.Context) {
	var body user.Dto
	if err := c.ShouldBind(&body); err != nil {
		c.HTML(http.StatusBadRequest, "error.html", gin.H{
			"Error": err.Error(),
		})
		return
	}
	if err := user.Create(body.Parse()); err != nil {
		c.HTML(http.StatusBadRequest, "error.html", gin.H{
			"Error": err.Error(),
		})
		return
	}
	c.String(http.StatusOK, "done")
	return
}

func RegisterPage(c *gin.Context){
	c.HTML(http.StatusOK, "user.html", gin.H{})
	return
}