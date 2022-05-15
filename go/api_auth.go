package main

// apiによるsignup usernameが重複している時は別のエラーを返すようにする
// func signup(w http.ResponseWriter, r *http.Request) {
// 	if r.Method == "GET" {
// 		_, err := session(w, r)
// 		if err != nil {
// 			generateHTML(w, nil, "layout", "public_navbar", "signup")
// 		} else {
// 			http.Redirect(w, r, "/posts", 302)
// 		}
// 	} else if r.Method == "POST" {
// 		err := r.ParseForm()
// 		if err != nil {
// 			log.Fatalln(err)
// 		}
// 		user := &User{
// 			Name:     r.PostFormValue("name"),
// 			PassWord: r.PostFormValue("password"),
// 		}
// 		if err := user.CreateUser(); err != nil {

// 			if err.(*mysql.MySQLError).Number == 1062 {
// 				log.Println("duplicate error")
// 				http.Redirect(w, r, "/signup", 302)
// 			} else {
// 				http.Redirect(w, r, "/signup", 302)
// 			}
// 		}
// 		http.Redirect(w, r, "/", 302)
// 	}
// }
