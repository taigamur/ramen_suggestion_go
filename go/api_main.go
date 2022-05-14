package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type ResponseUser struct {
	Status int    `json:"status"`
	Users  []User `json:"users"`
}

func getAllUsers(w http.ResponseWriter, r *http.Request) {
	users, _ := GetAllUsers()
	response := ResponseUser{http.StatusOK, users}
	res, _ := json.Marshal(response)

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, string(res))
}
