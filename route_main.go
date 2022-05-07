package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func top(w http.ResponseWriter, r *http.Request) {
	_, err := session(w, r)
	if err != nil {
		generateHTML(w, "welcome", "public_navbar", "layout", "top")
	} else {
		http.Redirect(w, r, "/posts", 302)
	}
}

func index(w http.ResponseWriter, r *http.Request) {
	sess, err := session(w, r)
	if err != nil {
		http.Redirect(w, r, "/login", 302)
	} else {
		user, err := sess.GetUserBySession()
		if err != nil {
			log.Println(err)
		}
		posts, _ := GetPosts(user.ID)
		user.Posts = posts
		generateHTML(w, user, "layout", "private_navbar", "index")
	}
}

func postNew(w http.ResponseWriter, r *http.Request) {
	_, err := session(w, r)
	if err != nil {
		http.Redirect(w, r, "/login", 302)
	} else {
		generateHTML(w, nil, "layout", "private_navbar", "post_new")
	}
}

func postSave(w http.ResponseWriter, r *http.Request) {
	sess, err := session(w, r)
	if err != nil {
		http.Redirect(w, r, "/login", 302)
	} else {
		err = r.ParseForm()
		if err != nil {
			log.Println(err)
		}
		user, err := sess.GetUserBySession()
		if err != nil {
			log.Println(err)
		}
		comment := r.PostFormValue("comment")
		place_id, _ := strconv.Atoi(r.PostFormValue("place_id"))
		value, _ := strconv.Atoi(r.PostFormValue("value"))
		if err := user.CreatePost(place_id, value, comment); err != nil {
			log.Println(err)
		}
		http.Redirect(w, r, "/posts", 302)
	}
}

func postEdit(w http.ResponseWriter, r *http.Request, id int) {
	sess, err := session(w, r)
	if err != nil {
		http.Redirect(w, r, "/login", 302)
	} else {
		_, err := sess.GetUserBySession()
		if err != nil {
			log.Println(err)
		}
		t, err := GetPost(id)
		if err != nil {
			log.Println(err)
		}
		generateHTML(w, t, "layout", "private_navbar", "post_edit")
	}
}

func postUpdate(w http.ResponseWriter, r *http.Request, id int) {
	sess, err := session(w, r)
	if err != nil {
		http.Redirect(w, r, "/login", 302)
	} else {
		_, err := sess.GetUserBySession()
		if err != nil {
			log.Println(err)
		}
		user, err := sess.GetUserBySession()
		if err != nil {
			log.Println(err)
		}
		comment := r.PostFormValue("comment")
		place_id, _ := strconv.Atoi(r.PostFormValue("place_id"))
		value, _ := strconv.Atoi(r.PostFormValue("value"))
		t := Post{ID: id, Comment: comment, PlaceID: place_id, Value: value, UserID: user.ID}
		if err := t.UpdatePost(); err != nil {
			log.Println(err)
		}
		http.Redirect(w, r, "/posts", 302)
	}
}

func postDelete(w http.ResponseWriter, r *http.Request, id int) {
	sess, err := session(w, r)
	if err != nil {
		http.Redirect(w, r, "/login", 302)
	} else {
		err = r.ParseForm()
		if err != nil {
			log.Println(err)
		}
		_, err := sess.GetUserBySession()
		if err != nil {
			log.Println(err)
		}
		t, err := GetPost(id)
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Println(t)
		if err := t.DeletePost(); err != nil {
			log.Println(err)
		}
		http.Redirect(w, r, "/posts", 302)
	}
}

func places(w http.ResponseWriter, r *http.Request) {
	sess, err := session(w, r)
	if err != nil {
		http.Redirect(w, r, "/login", 302)
	} else {
		user, err := sess.GetUserBySession()
		if err != nil {
			log.Println(err)
		}
		places, _ := GetPlaces()
		user.Places = places
		generateHTML(w, user, "layout", "private_navbar", "place_index")
	}
}
