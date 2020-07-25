package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Article struct {
	Title   string `json:"Title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

type Articles []Article

func getArticles(w http.ResponseWriter, r *http.Request) {

	articles := Articles{
		Article{Title: "Jon", Desc: "Haitohai", Content: "Heloooo"},
	}
	fmt.Println("Endpoint Hii: All Articles EndPoint")
	json.NewEncoder(w).Encode(articles)
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "homepage Endpoint hii")
}

func jonsyPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Jonsy Endpoint hii")
}

func postArticles(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "POST Request Checked")
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homePage).Methods("GET")
	myRouter.HandleFunc("/jonsy", jonsyPage).Methods("GET")
	myRouter.HandleFunc("/articles", getArticles).Methods("GET")
	myRouter.HandleFunc("/articles", postArticles).Methods("POST")
	log.Fatal(http.ListenAndServe(":8081", myRouter))
}

func main() {
	handleRequests()
}
