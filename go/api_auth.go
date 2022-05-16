package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/go-sql-driver/mysql"
)

// apiによるsignup usernameが重複している時は別のエラーを返すようにする
func signup(w http.ResponseWriter, r *http.Request) {
	setApiHeader(w)
	request := &User{
		Name:     r.PostFormValue("name"),
		PassWord: r.PostFormValue("password"),
	}
	if err := request.CreateUser(); err != nil {
		if err.(*mysql.MySQLError).Number == 1062 {
			w.WriteHeader(http.StatusForbidden)
			log.Println("duplicate error")
		} else {
			w.WriteHeader(http.StatusOK)
			apiuser, _ := GetApiUser(request.Name)
			res, _ := json.Marshal(apiuser)
			fmt.Fprint(w, string(res))
			log.Println("signup ok")
		}
	}
}

func login(w http.ResponseWriter, r *http.Request) {
	setApiHeader(w)
	request := &User{
		Name:     r.PostFormValue("name"),
		PassWord: r.PostFormValue("password"),
	}
	user, _ := GetUserByUserName(request.Name)
	log.Println(request.Name)
	log.Println(request.PassWord)

	if user.PassWord == Encrypt(request.PassWord) {
		w.WriteHeader(http.StatusOK)
		apiuser, _ := GetApiUser(request.Name)
		res, _ := json.Marshal(apiuser)
		fmt.Fprint(w, string(res))
	} else {
		w.WriteHeader(http.StatusForbidden)
	}

}
