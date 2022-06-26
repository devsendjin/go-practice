package main

import (
	"fmt"
	"net/http"

	"github.com/devsendjin/go-learn/pkg/handlers"
)

const appPort = ":8081";

// main in the main application function
func main()  {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Printf("Starting application on port %s\n\n\n", appPort);

	http.ListenAndServe(appPort, nil)
}
