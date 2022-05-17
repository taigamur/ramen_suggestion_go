package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func getAllUsers(w http.ResponseWriter, r *http.Request) {
	data, _ := GetAllUsers()
	res, _ := json.Marshal(data)

	setApiHeader(w)
	// w.WriteHeader(http.StatusForbidden)
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, string(res))
}

func postNew(w http.ResponseWriter, r *http.Request) {
	setApiHeader(w)
	// request := &User{
	// 	Name: r.PostFormValue("name"),
	// }

	// posts, _ := GetPosts(request.Name)
	w.WriteHeader(http.StatusOK)
}
