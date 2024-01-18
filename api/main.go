package main

import (
	"fmt"
	"net/http"
)

func globalHandler(w http.ResponseWriter, r *http.Request) {
	// handle HTTP requests
	r.URL.Path = "Hello Path!"
	fmt.Println(r.URL.Path)
	fmt.Fprintf(w, "Hello World!")
}

func main() {
	fmt.Println("Hello World!")

	globalHandler := http.HandlerFunc(globalHandler)

	http.ListenAndServe("127.0.0.1:8085", globalHandler)
}
