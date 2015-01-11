package handlers

import (
    "net/http"
    "fmt"
    
    "github.com/gorilla/mux"
)

func Redirect(w http.ResponseWriter, request *http.Request) {

	vars := mux.Vars(request)
	category := vars["hash"]
	
	fmt.Println(category)

}