package handlers

import (
    "fmt"
	"html/template"
    "net/http"
    netUrl "net/url"
    
    "github.com/mraxus/short-url/engine"
)

// Http handle that takes a url form parameter and converts it into a shortened URL
func Shorten(w http.ResponseWriter, r *http.Request) {
	
	filename := "templates/shorten.html"
	
    t, err := template.ParseFiles(filename)
    
    if err != nil {
    	fmt.Println("Error: Could not load template", filename)
		http.Error(w, http.StatusText(500), 500)
        return 
    }
    
    r.ParseForm()
    
    if r.Form["url"] == nil {
    	fmt.Println("form field 'url' is not given")
		http.Error(w, "form field 'url' is not given", 400)
		return
    }
    
    url := r.FormValue("url")
    _, err = netUrl.ParseRequestURI(url)
    
    if err != nil {
    	fmt.Println("form field 'url' is not correctly formatted")
		http.Error(w, "form field 'url' is not correctly formatted", 400)
		return
    }
    
    key := engine.Shorten(url)
    shortenedUrl := GetHost() + key
    
    fmt.Println(url + " --> " + shortenedUrl)
    
	t.Execute(w, map[string] string { "ShortenedUrl": shortenedUrl })
}
