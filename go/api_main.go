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
	fmt.Fprint(w, string(res))
}
