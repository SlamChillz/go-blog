package main

import (
	"net/http"
	"log"

	"github.com/gorilla/mux"
)

func main() {
	defer db.Close()
	router := mux.NewRouter()
	fileServer := http.FileServer(http.Dir("./static"))
	router.HandleFunc("/blog", PostList)
	router.HandleFunc("/blog/{id:[0-9]+}/{slug:[a-z-_0-9]+}", PostDetail)
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static", fileServer))
	router.Handle("/", http.RedirectHandler("/blog", http.StatusTemporaryRedirect))
	err := http.ListenAndServe(":5000", router)
	if err != nil {
		log.Fatal(err)
	}
}
