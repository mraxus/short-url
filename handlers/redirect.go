package handlers

import (
    "log"
    "net/http"
    
    "github.com/gorilla/mux"
    
    "github.com/mraxus/short-url/engine"
)

func Redirect(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	hash := vars["hash"]
	url, exists := engine.Resolve(hash)
	
	if exists {
		log.Println("Redirecting", hash, "-->", url)
		http.Redirect(w, r, url, 301)
		return
	}
	
	log.Println("Cannot redirect", hash, ": not found")
	http.Error(w, http.StatusText(404), 404)
}