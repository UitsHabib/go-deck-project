package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}

func handleRequests() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", helloWorld).Methods("GET")
	router.HandleFunc("/users", AllUsers).Methods("GET")
	router.HandleFunc("/users/{name}/{email}", NewUser).Methods("POST")
	router.HandleFunc("/users/{name}", DeleteUser).Methods("DELETE")
	router.HandleFunc("/users/{name}/{email}", UpdateUser).Methods("PUT")
	log.Fatal(http.ListenAndServe(":8081", router))
}

func main() {
	fmt.Println("Go ORM Tutorials")

	InitialMigration()

	handleRequests()
}
