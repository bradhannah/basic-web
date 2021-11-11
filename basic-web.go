package main

import (
	"fmt"
	"log"
	"net/http"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func accountHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Getting an account!!")
}
func accountCreate(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Creating an account!!")
}
func accountFetch(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Fetching an account!!")
}
func accountDelete(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Deleting an account!!")
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/api/account", accountHandler)
	http.HandleFunc("/api/account/create", accountCreate)
	http.HandleFunc("/api/account/fetch", accountFetch)
	http.HandleFunc("/api/account/delete", accountDelete)

	log.Fatal(http.ListenAndServe(":10000", nil))
}

func main() {
	handleRequests()
}
