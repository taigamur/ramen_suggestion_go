package main

import (
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"
	"sync"

	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/objx"
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
	data := map[string]interface{}{
		"Host": r.Host,
	}
	if authCookie, err := r.Cookie("auth"); err == nil {
		data["UserData"] = objx.MustFromBase64(authCookie.Value)
	}
	t.templ.Execute(w, data)
}

func main() {
	oauthSetup()

	// http.Handle("/", MustAuth(&templateHandler{filename: "top.html"}))
	// http.Handle("/login", &templateHandler{filename: "login.html"})
	// http.HandleFunc("/auth/", loginHandler)

	fmt.Println("TEST DB")

	user, _ := GetUser(1)
	user.CreatePost(1, 10, "okok")

	// log.Fatal(http.ListenAndServe(":8080", nil))

}
