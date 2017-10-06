package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func main() {
	config := new(Config)

	config.Port = os.Getenv("PORT")

	if config.Port == "" {
		config.Port = "8080"
	}

	config.Engine = gin.Default()
	config.DB, _ = sqlx.Open("sqlite3", ":memory:")

	config.InitTemplates(`templates`)
	config.InitStatic("/static", "static")

	InitHandlers(config)

	config.Engine.Run(":" + config.Port)
}
