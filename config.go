package main

import (
	"fmt"
	"html/template"
	"log"
	"os"

	rice "github.com/GeertJohan/go.rice"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

// Config for setting up gin ending and database
type Config struct {
	Engine *gin.Engine
	DB     *sqlx.DB
	Port   string
}

// InitTemplates initialize config templates
func (config *Config) InitTemplates(path string) {
	if path != "templates" {
		panic(fmt.Sprintf("Change rice.MustFindBox(`%s`) below.", path))
	}
	templateBox := rice.MustFindBox(`templates`)
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
	config.Engine.SetHTMLTemplate(templates)
}

// InitStatic initialize config static files
func (config *Config) InitStatic(urlPrefix string, folderPath string) {
	if folderPath != "static" {
		panic(fmt.Sprintf("Change rice.MustFindBox(`%s`) below.", folderPath))
	}
	staticBox := rice.MustFindBox(`static`)
	config.Engine.StaticFS(urlPrefix, staticBox.HTTPBox())
}
