package config

import (
	"html/template"
	"log"
)

// Appconfig hosds the application config
type AppConfig struct {
	Usecache      bool
	TemplateCache map[string]*template.Template
	InfoLog       *log.Logger
}
