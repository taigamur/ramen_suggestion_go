package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

// type templateHandler struct {
// 	once     sync.Once
// 	filename string
// 	templ    *template.Template
// }

// func (t *templateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
// 	t.once.Do(func() {
// 		t.templ =
// 			template.Must(template.ParseFiles(filepath.Join("templates", t.filename)))
// 	})
// 	t.templ.Execute(w, r)
// }

func generateHTML(w http.ResponseWriter, data interface{}, filenames ...string) {
	var files []string
	for _, file := range filenames {
		files = append(files, fmt.Sprintf("templates/%s.html", file))
	}
	templates := template.Must(template.ParseFiles(files...))
	templates.ExecuteTemplate(w, "layout", data)
}

func main() {

	http.HandleFunc("/", top)

	// user, _ := GetUser(1)
	// user.CreatePost(1, 10, "okok")

	log.Fatal(http.ListenAndServe(":8080", nil))

}
