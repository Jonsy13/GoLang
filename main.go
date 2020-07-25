package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
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
	fmt.Println("Endpoint Hii: All Articles EndPoint Here")
	json.NewEncoder(w).Encode(articles)
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "homepage Endpoint hii")
}

func jonsyPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Jonsy Endpoint hii")
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/jonsy", jonsyPage)
	http.HandleFunc("/articles", getArticles)
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func main() {
	handleRequests()
}
