package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/devsendjin/go-learn/pkg/config"
	"github.com/devsendjin/go-learn/pkg/handlers"
	"github.com/devsendjin/go-learn/pkg/render"
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

	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Printf("Starting application on port %s\n\n\n", appPort)

	http.ListenAndServe(appPort, nil)
}
