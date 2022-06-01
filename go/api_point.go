package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func pointIndex(w http.ResponseWriter, r *http.Request) {
	setApiHeader(w)
	username := r.URL.Query().Get("username")
	points := GetPoints(username)

	w.WriteHeader(http.StatusOK)
	res, _ := json.Marshal(points)
	fmt.Fprint(w, string(res))

}

func pointNew(w http.ResponseWriter, r *http.Request) {
	setApiHeader(w)
	username := r.PostFormValue("username")
	place_id, _ := strconv.Atoi(r.PostFormValue("place_id"))
	value, _ := strconv.Atoi(r.PostFormValue("value"))

	err := CreatePoint(username, place_id, value)

	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
	} else {
		w.WriteHeader(http.StatusOK)
	}

}

func pointUpdate(w http.ResponseWriter, r *http.Request) {
	setApiHeader(w)
	username := r.PostFormValue("username")
	place_id, _ := strconv.Atoi(r.PostFormValue("place_id"))
	value, _ := strconv.Atoi(r.PostFormValue("value"))

	err := UpdatePoint(username, place_id, value)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	} else {
		w.WriteHeader(http.StatusOK)
	}
}
