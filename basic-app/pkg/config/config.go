package config

import "html/template"

/*
We can store all the configuration for our application in this struct
*/
type AppConfig struct {
	TemplateCache map[string]*template.Template
}
