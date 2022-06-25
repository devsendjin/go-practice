package main

import (
	"fmt"
	"net/http"
)

const appPort = ":8081";

// main in the main application function
func main()  {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	fmt.Printf("Starting application on port %s\n", appPort);

	http.ListenAndServe(appPort, nil)
}
