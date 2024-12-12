package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/raywu60kg/miniMD/internal/handlers"
)

func main() {
	// User handler routes
	http.HandleFunc("/", handlers.GetHome)

	// Start server
	port := 2486
	fmt.Printf("Server starting on port %d\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}
