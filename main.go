package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Post - Our struct for all articles
type Post struct {
	ID      int    `json:"ID"`
	Title   string `json:"Title"`
	Content string `json:"content"`
}

// Posts : An array of artcles
type Posts []Post

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Home ? VueJS Frontend?")
}

func returnPosts(w http.ResponseWriter, r *http.Request) {
	posts := Posts{
		Post{Title: "Post Title One", Content: "This post contains some not so interesting data."},
		Post{Title: "Post Title Two", Content: "This post contains some not so interesting data."},
	}

	json.NewEncoder(w).Encode(posts)
}

func returnPost(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]
	fmt.Fprintf(w, "ID : "+key)
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", home)
	myRouter.HandleFunc("/all", returnPosts)
	myRouter.HandleFunc("/post/{id}", returnPost)
	log.Fatal(http.ListenAndServe(":8081", myRouter))
}

func main() {
	handleRequests()
}
