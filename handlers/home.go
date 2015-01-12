package handlers

import (
    "fmt"
	"io/ioutil"
    "log"
    "net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	
	filename := "templates/index.html"
	
    body, err := ioutil.ReadFile(filename)
    
    if err != nil {
    	log.Println("Error: Could not load html file", filename)
		http.Error(w, http.StatusText(500), 500)
        return 
    }

	fmt.Fprintf(w, "%s", body)
}
