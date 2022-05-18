package main

import (
	"log"
	"net/http"
	"strconv"
)

func postNew(w http.ResponseWriter, r *http.Request) {
	setApiHeader(w)

	place_id, _ := strconv.Atoi(r.PostFormValue("place_id"))
	value, _ := strconv.Atoi(r.PostFormValue("value"))
	username := r.PostFormValue("username")
	date := r.PostFormValue("date")

	err := CreatePost(place_id, username, value, date)

	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		log.Println(err)
	} else {
		w.WriteHeader(http.StatusOK)
	}
}
