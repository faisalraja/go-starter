package main

// As this grows you can move it to it's own package same with config and your models

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// InitHandlers create routes
func InitHandlers(config *Config) {
	config.Engine.GET("/", home)

	config.Engine.NoRoute(func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl.html", nil)
	})
}

func home(c *gin.Context) {
	c.HTML(http.StatusOK, "index.tmpl.html", nil)
}
