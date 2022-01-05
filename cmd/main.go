package main

import (
	"currency-converter/pkg/controllers"
	"currency-converter/pkg/model"
	"flag"
	"fmt"
	"net/http"
)

// Application entry point
func main() {
	// parse command line argument
	domain := flag.String("domain", "localhost:8080", "host and port to run service on")
	flag.Parse()
	// configure database
	if err := model.SetupDatabase(); err != nil {
		fmt.Printf("database error: %+v", err)
		panic(err)
	}
	// configure request handler
	http.HandleFunc("/convert-ccy", controllers.GetExcahngeRate)
	fmt.Printf("Starting HTTP server @ %+v\n", *domain)
	// listen for requests
	if err := http.ListenAndServe(*domain, nil); err != nil {
		fmt.Printf("Error starting server: %v\n", err)
		panic(err)
	}
}
