package main

import (
	"fmt"
	"net/http"
)

func main()  {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		n, err := fmt.Fprintf(w, "Hello world!")

		if err != nil {
			 fmt.Println(err)
		}

		fmt.Println(fmt.Sprintf("NUmber of bytes written: %d", n))
	})

	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		fmt.Println(err)
	}
}
