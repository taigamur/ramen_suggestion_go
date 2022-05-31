package main

import (
	"fmt"
	"log"
	"net/http"
)

func test(w http.ResponseWriter, r *http.Request) {
	setApiHeader(w)
	log.Println("test")
	fmt.Fprint(w, "test")
}
