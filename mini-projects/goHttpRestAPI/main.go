package main

import (
	"fmt"
	"goHttpRestAPI/handlers"
	"net/http"
)

/*
	func rootHandler(w http.ResponseWriter, r *http.Request) {
		// w is interface to response writer and r is for http Request
		if r.URL.Path != "/" {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("Requested URL not found i.e. 404 \n"))
			return
		}
		w.WriteHeader(http.StatusOK)                       //instead of code use constant
		w.Write([]byte("API Version v1 : Hello World \n")) //Simple conversion
	}
*/
func main() {
	fmt.Println("Starting go web server")
	//http.HandleFunc("/", rootHandler)
	http.HandleFunc("/", handlers.RootHandler)
	http.HandleFunc("/users", handlers.UsersRouter)
	http.HandleFunc("/users/", handlers.UsersRouter)
	err := http.ListenAndServe("localhost:8086", nil)
	if err != nil {
		fmt.Printf("Encountered an error while starting server %s", err)
	}
}

/*
API Design Spec before starting development

Resource /user
List item data, id, name , age, location, YOE (Year of Experience), role , IC - true/false, EM - true/false, Exec - true/false
Functionality - CRUD
endpoints
	- /user/
		- /user/{id}
  	- /role/{exec,em,ic}
		- Return user which above queries roles

Data Formatting
 	- JSON
*/
