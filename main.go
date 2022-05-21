package main

// init my imports
import (
	"fmt"
	"log"
	"net/http"
	"simple_web_server/handlers"
)

func main() {
	// make simple server over /static directory
	fileServer := http.FileServer(http.Dir("./static"))

	// implement our routes
	http.Handle("/", fileServer)
	http.HandleFunc("/form", handlers.FormHandler)
	http.HandleFunc("/hello", handlers.HelloHandler)

	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
