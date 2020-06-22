package main

import (
	"fmt"
	_ "github.com/joho/godotenv/autoload"
	"github.com/labstack/echo/v4"
	"github.com/sorushsaghari/ie/cmd/html/internal/note"
	"github.com/sorushsaghari/ie/cmd/html/internal/user"
	n "github.com/sorushsaghari/ie/internal/note"
	"github.com/sorushsaghari/ie/internal/platforms/cfg"
	"github.com/sorushsaghari/ie/internal/platforms/database"
	u "github.com/sorushsaghari/ie/internal/user"
	"html/template"
	"io"
	"log"
)

type TemplateRenderer struct {
	templates *template.Template
}

// Render renders a template document
func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {

	// Add global methods if data is a map
	if viewContext, isMap := data.(map[string]interface{}); isMap {
		viewContext["reverse"] = c.Echo().Reverse
	}

	return t.templates.ExecuteTemplate(w, name, data)
}

func main(){
	cfg.Load()
	err := database.Init(cfg.GetConfig().DB_PASSWORD,
		cfg.GetConfig().DB_USER,
		cfg.GetConfig().DB_PORT,
		cfg.GetConfig().DB_NAME,
		cfg.GetConfig().DB_HOST)
	if err != nil {
		log.Fatal(err.Error())
	}
	database.DB().AutoMigrate(u.User{}, u.Auth{}, n.Note{})
	router := echo.New()
	renderer := &TemplateRenderer{
		templates: template.Must(template.ParseGlob("templates/**/*.html")),
	}
	router.Renderer = renderer
	user.Routers(router)
	note.Routers(router)
	e := router.Start(":8000")
	fmt.Println(e)
	router.Logger.Fatal(e)
}
