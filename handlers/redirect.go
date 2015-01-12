package handlers

import (
    "fmt"
    "net/http"
    
    "github.com/gorilla/mux"
    
    "github.com/mraxus/short-url/engine"
)

func Redirect(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	hash := vars["hash"]
	url, exists := engine.Resolve(hash)
	
	if exists {
		fmt.Println("Redirecting", hash, "-->", url)
		http.Redirect(w, r, url, 301)
		return
	}
	
	fmt.Println("Cannot redirect", hash, ": not found")
	http.Error(w, http.StatusText(404), 404)
}