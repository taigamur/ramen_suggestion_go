package main

import (
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"
	"sync"

	_ "github.com/go-sql-driver/mysql"
)

type templateHandler struct {
	once     sync.Once
	filename string
	templ    *template.Template
}

func (t *templateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	t.once.Do(func() {
		t.templ =
			template.Must(template.ParseFiles(filepath.Join("templates", t.filename)))
	})
	t.templ.Execute(w, r)
}

func main() {

	http.Handle("/", &templateHandler{filename: "top.html"})

	fmt.Println("TEST DB")

	user, _ := GetUser(1)
	user.CreatePost(1, 10, "okok")

	// log.Fatal(http.ListenAndServe(":8080", nil))

}
