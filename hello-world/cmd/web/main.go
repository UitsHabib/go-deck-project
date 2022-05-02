package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/UitsHabib/basic-web/pkg/config"
	"github.com/UitsHabib/basic-web/pkg/handlers"
	"github.com/UitsHabib/basic-web/pkg/render"
)

const portNumber = ":5050"

func main() {
	var app config.AppConfig

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("Cannot create template cache")
	}

	app.TemplateCache = tc
	app.Usecache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))

	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	log.Fatal(err)
}
