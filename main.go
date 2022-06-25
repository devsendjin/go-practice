package main

import (
	"errors"
	"fmt"
	"net/http"
)

const appPort = ":8081";

// Home is the home page handler
func Home(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintf(w, "This is the home page")
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

// Divide is the divide page handler
func Divide(w http.ResponseWriter, r *http.Request) {
	x := float32(100.0)
	y := float32(10.0) // change to 0 to throw error

	f, err := divideValues(x, y)
	if err != nil {
		fmt.Fprintf(w, "Cannot divide by 0")
		return
	}

	fmt.Fprintf(w, "%f divided by %f is %f", x, y, f)
}

func divideValues(x, y float32) (float32, error) {
	if y <=0 {
		err := errors.New("can not divide by zero")
		return 0, err
	}

	return x / y, nil
}

// main in the main application function
func main()  {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)
	http.HandleFunc("/divide", Divide)

	fmt.Printf("Starting application on port %s\n", appPort);

	http.ListenAndServe(appPort, nil)
}
