package main

import (
	"log"
	"time"
)

type User struct {
	Name      string
	PassWord  string
	CreatedAt time.Time
	Posts     []Post
	Places    []Place
}

type apiUser struct {
	Name string `json:"name"`
}

func (u *User) CreateUser() (err error) {
	cmd := `insert into users (
		name,
		password) values (?, ?)`

	_, err = Db.Exec(cmd,
		u.Name,
		Encrypt(u.PassWord))

	if err != nil {
		log.Println(err)
	}
	return err
}

func GetUser(id int) (user User, err error) {
	user = User{}
	cmd := `select name, password, created_at from users where id = ?`
	err = Db.QueryRow(cmd, id).Scan(
		&user.Name,
		&user.PassWord,
		&user.CreatedAt,
	)
	return user, err
}

func GetApiUser(name string) (user apiUser, err error) {
	user = apiUser{}
	cmd := `select name from users where name = ?`
	err = Db.QueryRow(cmd, name).Scan(
		&user.Name,
	)
	return user, err
}

func GetUserByUserName(user_name string) (user User, err error) {
	user = User{}
	cmd := `select name, password, created_at from users where name = ?`
	err = Db.QueryRow(cmd, user_name).Scan(
		&user.Name,
		&user.PassWord,
		&user.CreatedAt,
	)
	return user, err
}
