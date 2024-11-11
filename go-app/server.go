package main

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", Hello)

	fmt.Printf("Server starting on 8080 port \n")

	err := http.ListenAndServe(":8080", nil)
  if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("Server closed\n")
	} else if err != nil {
		fmt.Printf("Error starting server: %s\n", err)
		os.Exit(1)
	}
}

func Hello(w http.ResponseWriter, r *http.Request) {
	
	name := os.Getenv("NAME")
	age := os.Getenv("AGE")
	
	io.WriteString(w, "This is my website v4!\n")
	fmt.Fprintf(w, "Hello, I'm %s. I'm %s", name, age)
}