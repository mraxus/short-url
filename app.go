package main 

import (
    "net/http"
    "fmt"
    "os"
    
    "github.com/gorilla/mux"
    "github.com/mraxus/short-url/handlers"
)

func main() {
	
	var port string = os.Getenv("PORT")
	
	if port == "" { port = "8080" }
	

	registerRoutes()

    
    fmt.Println("String server on port " + port)
    err := http.ListenAndServe(":" + port, nil)
    if err != nil {
      panic(err)
    }
}


func registerRoutes() {
	
    r := mux.NewRouter()
    
    // Main page, shows URL generator form 
    r.HandleFunc("/", handlers.Home).Methods("GET")
    
    // Main page, shows URL generator form
    r.HandleFunc("/shorten", handlers.Shorten).Methods("POST")
    
    // Request to be redirected to original URL
    r.HandleFunc("/{hash}", handlers.Redirect).Methods("GET")
	
	http.Handle("/", r)
}