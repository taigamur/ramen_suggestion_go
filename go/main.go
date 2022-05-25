package main

import (
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	http.HandleFunc("/signup", signup)
	http.HandleFunc("/login", login)
	http.HandleFunc("/post/index", postIndex)
	http.HandleFunc("/post/new", postNew)
	http.HandleFunc("/place/index", getPlaces)
	http.HandleFunc("/place/suggest", suggestPlace)

	log.Fatal(http.ListenAndServe(":8080", nil))

}
