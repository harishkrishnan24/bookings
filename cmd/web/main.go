package main

import (
	"encoding/gob"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/harishkrishnan24/bookings/internal/config"
	"github.com/harishkrishnan24/bookings/internal/handlers"
	"github.com/harishkrishnan24/bookings/internal/models"
	"github.com/harishkrishnan24/bookings/internal/render"
)

const portNumber = ":8080"

var app config.AppConfig
var session *scs.SessionManager

func main() {
	err := run()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))

	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}

func run() error {
	gob.Register(models.Reservation{})

	app.InProduction = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
		return err
	}

	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)
	return nil
}
