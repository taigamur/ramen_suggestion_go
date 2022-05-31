package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func postNew(w http.ResponseWriter, r *http.Request) {
	setApiHeader(w)

	place_id, _ := strconv.Atoi(r.PostFormValue("place_id"))
	value, _ := strconv.Atoi(r.PostFormValue("point"))
	username := r.PostFormValue("username")
	date := r.PostFormValue("date")

	err := CreatePost(place_id, username, value, date)

	// pointの追加
	point := GetPoint(username, place_id)
	if point.UserName == "" { //point tableにないとき
		fmt.Println("create point")
		CreatePoint(username, place_id, value)
	} else {
		fmt.Println("update point")
		UpdatePoint(username, place_id, value)
	}

	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		log.Println(err)
	} else {
		w.WriteHeader(http.StatusOK)
	}
}

func postIndex(w http.ResponseWriter, r *http.Request) {
	setApiHeader(w)

	username := r.URL.Query().Get("username")

	posts, err := GetPosts(username)

	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		log.Println(err)
	} else {
		w.WriteHeader(http.StatusOK)
		res, _ := json.Marshal(posts)
		fmt.Fprint(w, string(res))
	}
}

func getPlaces(w http.ResponseWriter, r *http.Request) {
	setApiHeader(w)
	keyword := r.PostFormValue("keyword")
	places, err := GetPlacesByKeyword(keyword)

	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		log.Println(err)
	} else {
		if keyword == "" || len(places) == 0 {
			w.WriteHeader(http.StatusNotFound)
			log.Println("404: Not Found")
		} else {
			w.WriteHeader(http.StatusOK)
			res, _ := json.Marshal(places)
			fmt.Fprint(w, string(res))
		}
	}
}

func postDelete(w http.ResponseWriter, r *http.Request) {
	setApiHeader(w)
	id, _ := strconv.Atoi(r.PostFormValue("id"))
	username := r.PostFormValue("username")

	post, _ := GetPost(id)
	if post.UserName == username {
		err := DeletePost(id)
		if err != nil {
			log.Println("post delete error")
			w.WriteHeader(http.StatusBadRequest)
		} else {
			w.WriteHeader(http.StatusOK)
		}
	} else {
		w.WriteHeader(http.StatusBadRequest)
	}
}
