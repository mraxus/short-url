package handlers

import (
    "fmt"
	"html/template"
    "net/http"
)

func Shorten(w http.ResponseWriter, r *http.Request) {
	
	filename := "templates/shorten.html"
	
    t, err := template.ParseFiles(filename)
    
    if err != nil {
    	fmt.Fprintf(w, "Hi there, I could not load the page =(")
        return 
    }

	t.Execute(w, map[string] string {"ShortenedUrl": "http://www.short.com"})
}
