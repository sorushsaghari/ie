package user

import (
	"github.com/gin-gonic/gin"
	"github.com/sorushsaghari/ie/internal/user"
	"net/http"
)

func Register(c *gin.Context) {
	var body user.CreateDto
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

func LoginPage(c *gin.Context){
	c.HTML(http.StatusOK, "login.html", nil)
	return
}
func Login(c *gin.Context){
	var body user.Dto
	if err := c.ShouldBind(&body); err != nil {
		c.HTML(http.StatusBadRequest, "error.html", gin.H{
			"Error": err.Error(),
		})
		return
	}
	u, err := user.One(body.Parse())
	if err != nil || u == nil {
		c.HTML(http.StatusBadRequest, "error.html", gin.H{
			"Error": err.Error(),
		})
		return
	}
	var auth user.Auth
	auth.Token = user.TokenGenerator()
	auth.UserID = u.ID
	if err = user.Store(&auth);err != nil {
		c.HTML(http.StatusBadRequest, "error.html", gin.H{
			"Error": err.Error(),
		})
		return
	}
	c.SetCookie("token", auth.Token, 3600, "*", "127.0.0.1", false, true)
	c.String(http.StatusOK, "log in was successful")
	return
}
