package main 

import (
    "net/http"
    "fmt"
    "os" 
)

// Default Request Handler
func defaultHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "<h1>Hello %s!</h1>", r.URL.Path[1:])
}

// SubRequest Handler
func subHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "<h1>Heroko %s!</h1>", r.URL.Path[1:])
}

func main() {
	
	var port string = os.Getenv("PORT")
	
	if port == "" { port = "8080" }
	
    http.HandleFunc("/", defaultHandler)
    http.HandleFunc("/bye/", subHandler)
    
    fmt.Println("String server on port " + port)
    err := http.ListenAndServe(":" + port, nil)
    if err != nil {
      panic(err)
    }
}
