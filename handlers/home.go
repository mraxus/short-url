package handlers

import (
    "fmt"
	"io/ioutil"
    "net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	
	filename := "templates/index.html"
	
    body, err := ioutil.ReadFile(filename)
    
    if err != nil {
    	fmt.Fprintf(w, "Hi there, I could not load the page =(")
        return 
    }

	fmt.Fprintf(w, "%s", body)
}
