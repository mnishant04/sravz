package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

// Structure of type Request
type Request struct {
	Message string `json:"message"`
}

//Function for Getting /hello
func HelloGet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode("hello")
}

//Function for Posting String to /hello
func HelloPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	body, _ := ioutil.ReadAll(r.Body)
	var request Request
	json.Unmarshal(body, &request)
	json.NewEncoder(w).Encode("hello " + request.Message)
}

//Main Function
func main() {

	//Creating router
	r := mux.NewRouter()
	fmt.Println("Starting Server ")
	fmt.Println("Hosting Route")
	r.HandleFunc("/hello", HelloGet).Methods("GET")
	r.HandleFunc("/hello", HelloPost).Methods("POST")
	fmt.Println("Server started")
	http.ListenAndServe(":8080", r)
}
