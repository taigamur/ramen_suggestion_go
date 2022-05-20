package main

import (
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	http.HandleFunc("/signup", signup)
	http.HandleFunc("/login", login)
	// http.HandleFunc("/authenticate", authecticate)
	// http.HandleFunc("/logout", logout)
	http.HandleFunc("/post/index", postIndex)
	http.HandleFunc("/post/new", postNew)
	// http.HandleFunc("/posts/save", postSave)
	// http.HandleFunc("/posts/edit/", parseURL(postEdit))
	// http.HandleFunc("/posts/update/", parseURL(postUpdate))
	// http.HandleFunc("/posts/delete/", parseURL(postDelete))

	// http.HandleFunc("/places/index", places)

	log.Fatal(http.ListenAndServe(":8080", nil))

}
