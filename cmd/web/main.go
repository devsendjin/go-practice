package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/devsendjin/go-practice/pkg/config"
	"github.com/devsendjin/go-practice/pkg/handlers"
	"github.com/devsendjin/go-practice/pkg/render"
)

const appPort = ":8081"

// main in the main application function
func main() {
	var app config.AppConfig

	templateCache, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("Cannot create template cache")
	}

	app.TemplateCache = templateCache
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	fmt.Printf("\nStarting application on port %s\n\n\n", appPort)

	server := &http.Server{
		Addr:    appPort,
		Handler: routes(&app),
	}

	err = server.ListenAndServe()
	log.Fatal(err)
}
