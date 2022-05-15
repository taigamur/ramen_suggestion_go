package main

import (
	"log"
	"net/http"
)

func signup(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {
		err := r.ParseForm()
		if err != nil {
			log.Fatalln(err)
		}
		user := &User{
			Name:     r.PostFormValue("name"),
			Email:    r.PostFormValue("email"),
			PassWord: r.PostFormValue("password"),
		}
		if err := user.CreateUser(); err != nil {
			http.Redirect(w, r, "/signup", 302)
		}
		http.Redirect(w, r, "/", 302)
	}
}
