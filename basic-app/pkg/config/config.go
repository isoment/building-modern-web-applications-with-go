package config

import (
	"html/template"
	"log"
)

/*
We can store all the configuration for our application in this struct
*/
type AppConfig struct {
	UseCache      bool
	TemplateCache map[string]*template.Template
	InfoLog       *log.Logger
}
