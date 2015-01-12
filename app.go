package main 

import (
    "fmt"
	"math/rand"
    "net/http"
    "os"
	"time"
    
    "github.com/gorilla/mux"
    "github.com/mraxus/short-url/handlers"
)

// To randomize the key generator, making the generator non-deterministic
func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

// Starting app as a webserver with env params:
//   HOST - a complete http/https host URL with trailing slash and PORT if not default port ("http://short.url.nu:8080/")
//          Note that the running http server might listen to one port but the endpoint to the server might differ.
//
//   POST - which port the http server should listen to 
func main() {
	
	host := os.Getenv("HOST")
	port := os.Getenv("PORT")
	
	// Default to port:
	if port == "" { port = "8080" }
	
	// Set the hostname provided by the URL shortening service
	if host == "" {
		fmt.Println("No HOST env variable given; setting localhost as host on port " + port)
		handlers.SetHost("http://localhost:" + port + "/")
	} else {
		fmt.Println("Setting host: " + host)
		handlers.SetHost(host)
	}
	
	registerRoutes()

    // (Try to) Start the web server
    fmt.Println("Starting server on port " + port)
    err := http.ListenAndServe(":" + port, nil)
    if err != nil {
		panic(err)
    }
}

// Register all routes used by the web server
func registerRoutes() {
	
    r := mux.NewRouter()
    
    // Main page, shows URL generator form 
    r.HandleFunc("/", handlers.Home).Methods("GET")
    
    // Generator page, shows generated shortened URL
    r.HandleFunc("/shorten", handlers.Shorten).Methods("POST")
    
    // Request to be redirected to original URL
    r.HandleFunc("/{hash}", handlers.Redirect).Methods("GET")
	
	http.Handle("/", r)
}