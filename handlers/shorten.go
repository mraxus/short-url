package handlers

import (
	"html/template"
    "log"
    "net/http"
    
    "github.com/mraxus/short-url/engine"
)

var host = "http://localhost:8080/"

func Shorten(w http.ResponseWriter, r *http.Request) {
	
	filename := "templates/shorten.html"
	
    t, err := template.ParseFiles(filename)
    
    if err != nil {
    	log.Println("Error: Could not load template", filename)
		http.Error(w, http.StatusText(500), 500)
        return 
    }
    
    r.ParseForm()
    
    if r.Form["url"] == nil {
    	log.Println("form field 'url' is not given")
		http.Error(w, http.StatusText(400), 400)
		return
    }
    
    url := r.FormValue("url")
    
    log.Println("url", url)
    
    hash := engine.Shorten(url)
    
	t.Execute(w, map[string] string {"ShortenedUrl": host + hash})
}
