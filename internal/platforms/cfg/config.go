package cfg

import (
	"os"
	"strconv"
)

type Config struct {
	DB_NAME string 
	DB_USER string
	DB_PASSWORD string
	DB_HOST string
	DB_PORT int
}
var cfg Config
func Load()  {
	DB_PORT, _ := strconv.Atoi(os.Getenv("DB_PORT"))
	cfg = Config{
		DB_NAME:     os.Getenv("DB_NAME"),
		DB_USER:     os.Getenv("DB_USER"),
		DB_PASSWORD: os.Getenv("DB_PASSWORD"),
		DB_HOST:     os.Getenv("DB_HOST"),
		DB_PORT:    DB_PORT,
	}
}

func GetConfig() Config {
	return cfg
}