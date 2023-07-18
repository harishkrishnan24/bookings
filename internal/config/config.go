package config

import (
	"html/template"
	"log"

	"github.com/alexedwards/scs/v2"
	"github.com/harishkrishnan24/bookings/internal/models"
)

type AppConfig struct {
	UseCache      bool
	TemplateCache map[string]*template.Template
	InProduction  bool
	Session       *scs.SessionManager
	InfoLog       *log.Logger
	ErrorLog      *log.Logger
	MailChan      chan models.MailData
}
