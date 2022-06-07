package main

import (
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// auth
	http.HandleFunc("/signup", signup)
	http.HandleFunc("/login", login)
	// user
	http.HandleFunc("/user/id", getUserId)
	// post
	http.HandleFunc("/post/index", postIndex)
	http.HandleFunc("/post/new", postNew)
	http.HandleFunc("/post/delete", postDelete)
	// place
	http.HandleFunc("/place/index", getPlaces)
	http.HandleFunc("/place/suggest", suggestPlace)
	// point
	http.HandleFunc("/point/index", pointIndex)
	http.HandleFunc("/point/new", pointNew)
	http.HandleFunc("/point/update", pointUpdate)
	// public
	http.HandleFunc("/public/point/index", publicPointIndex)

	log.Fatal(http.ListenAndServe(":8080", nil))

}
