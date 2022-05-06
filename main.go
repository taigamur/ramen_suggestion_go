package main

import (
	"html/template"
	"log"
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
	t.templ.Execute(w, nil)
}

func main() {
	// db := connectDB()
	// defer db.Close()

	oauthSetup()

	http.Handle("/", MustAuth(&templateHandler{filename: "top.html"}))
	http.Handle("/login", &templateHandler{filename: "login.html"})
	http.HandleFunc("/auth/", loginHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
