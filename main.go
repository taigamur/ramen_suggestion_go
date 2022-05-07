package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

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
	http.HandleFunc("/signup", signup)
	http.HandleFunc("/login", login)
	http.HandleFunc("/authenticate", authecticate)

	user, _ := GetUser(1)
	user.CreatePost(1, 10, "okok")
	session1, _ := user.CreateSession()
	valid, _ := session1.CheckSession()
	fmt.Println(valid)

	user2, _ := GetUser(3)
	user2.CreateSession()

	session1.Email = "aaa.example.com"
	valid2, _ := session1.CheckSession()
	fmt.Println(valid2)

	log.Fatal(http.ListenAndServe(":8080", nil))

}
