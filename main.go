package main

import (
	"html/template"
	"log"
	"net/http"
	"os"

	rice "github.com/GeertJohan/go.rice"
	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	templateBox := rice.MustFindBox("templates")
	staticBox := rice.MustFindBox("static")

	router := gin.New()
	router.Use(gin.Logger())

	// Load all templates
	templates := template.New("Server Templates")
	templateBox.Walk("", func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		log.Print("Loading template: " + path)
		content, err := templateBox.String(path)
		if err != nil {
			return err
		}
		if _, err = templates.New(path).Parse(content); err != nil {
			return err
		}
		return nil
	})
	router.SetHTMLTemplate(templates)
	router.StaticFS("/static", staticBox.HTTPBox())
	

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl.html", nil)
	})

	router.Run(":" + port)
}
