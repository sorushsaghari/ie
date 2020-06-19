package note

import (
	"github.com/gin-gonic/gin"
	"github.com/sorushsaghari/ie/cmd/html/internal/middleware"
)

func Routers(r *gin.Engine) {
	note := r.Group("/note")
	note.Use(middleware.IsAuthenticated())
	note.GET("/", Index)
	//note.GET("/:id", Detail)
}
