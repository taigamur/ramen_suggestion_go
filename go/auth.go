package main

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

// func login(w http.ResponseWriter, r *http.Request) {
// 	_, err := session(w, r)
// 	if err != nil {
// 		generateHTML(w, nil, "layout", "public_navbar", "login")
// 	} else {
// 		http.Redirect(w, r, "/todos", 302)
// 	}
// }

// func authecticate(w http.ResponseWriter, r *http.Request) {
// 	err := r.ParseForm()
// 	user, err := GetUserByUserName(r.PostFormValue("name"))
// 	if err != nil {
// 		log.Println(err)
// 		http.Redirect(w, r, "/login", 302)
// 	}
// 	if user.PassWord == Encrypt(r.PostFormValue("password")) {
// 		session, err := user.CreateSession()
// 		if err != nil {
// 			log.Println(err)
// 		}
// 		cookie := http.Cookie{
// 			Name:     "_cookie",
// 			Value:    session.UserName,
// 			HttpOnly: true,
// 		}
// 		http.SetCookie(w, &cookie)
// 		http.Redirect(w, r, "/", 302)
// 	} else {
// 		http.Redirect(w, r, "/login", 302)
// 	}
// }

// func logout(w http.ResponseWriter, r *http.Request) {
// 	cookie, err := r.Cookie("_cookie")
// 	if err != nil {
// 		log.Fatalln(err)
// 	}
// 	if err != http.ErrNoCookie {
// 		session := Session{UserName: cookie.Value}
// 		session.DeleteSession()
// 	}
// 	http.Redirect(w, r, "/login", 302)
// }
