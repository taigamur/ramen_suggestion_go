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
