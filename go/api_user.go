package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func getUserId(w http.ResponseWriter, r *http.Request) {
	setApiHeader(w)
	username := r.URL.Query().Get("username")
	user_id, err := GetUserId(username)

	if err != nil {
		log.Println(err)
	}

	w.WriteHeader(http.StatusOK)
	res, _ := json.Marshal(user_id)
	fmt.Fprint(w, string(res))
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	setApiHeader(w)
	username := r.URL.Query().Get("username")
	err := DeleteUser(username)
	if err != nil {
		log.Println(err)
	}
	w.WriteHeader(http.StatusOK)
}
