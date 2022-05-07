package main

import "net/http"

func top(w http.ResponseWriter, r *http.Request) {
	generateHTML(w, "welcome", "public_navbar", "layout", "top")
}
