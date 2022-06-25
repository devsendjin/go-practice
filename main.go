package main

import (
	"fmt"
	"net/http"
)

const appPort = ":8081";

// Home is the home page handler
func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is the home page")
}

// About is the home page handler
func About(w http.ResponseWriter, r *http.Request) {
	sum := addValues(2, 2)
	_, _ = fmt.Fprintf(w, "This is the home page and 2 + 2 is %d", sum)
}

// addValues adds two integers and return the sum
func addValues(x, y int) int {
	return x + y
}

// main in the main application function
func main()  {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	fmt.Printf("Starting application on port %s", appPort);

	http.ListenAndServe(appPort, nil)
}
