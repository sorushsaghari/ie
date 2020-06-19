package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
	"github.com/sorushsaghari/ie/cmd/html/internal/note"
	"github.com/sorushsaghari/ie/cmd/html/internal/user"
	n "github.com/sorushsaghari/ie/internal/note"
	"github.com/sorushsaghari/ie/internal/platforms/cfg"
	"github.com/sorushsaghari/ie/internal/platforms/database"
	u "github.com/sorushsaghari/ie/internal/user"
	"log"
)

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
	router := gin.Default()
	router.LoadHTMLGlob("templates/**/*")
	user.Routers(router)
	note.Routers(router)
	router.Run()
}
