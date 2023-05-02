package config

import (
	"html/template"
	"log"

	"github.com/alexedwards/scs/v2"
)

/*
We can store all the configuration for our application in this struct
*/
type AppConfig struct {
	UseCache      bool
	TemplateCache map[string]*template.Template
	InfoLog       *log.Logger
	ErrorLog      *log.Logger
	InProduction  bool
	Session       *scs.SessionManager
}
