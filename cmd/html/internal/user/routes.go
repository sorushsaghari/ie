package user

import (
	"github.com/gin-gonic/gin"
	"github.com/sorushsaghari/ie/cmd/html/internal/middleware"
)

func Routers(r *gin.Engine) {
	user := r.Group("/user")
	user.GET("/", RegisterPage)
	user.POST("/", Register)
	user.GET("/login", LoginPage)
	user.POST("/login", Login)
	user.Use(middleware.IsAuthenticated())
}
