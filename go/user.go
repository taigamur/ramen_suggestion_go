package main

import (
	"log"
	"time"
)

type User struct {
	ID        int
	Name      string
	PassWord  string
	CreatedAt time.Time
	Posts     []Post
	Places    []Place
}

type apiUser struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Session struct {
	ID        int
	UserID    int
	UserName  string
	CreatedAt time.Time
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

// React 連携確認用のテスト
func GetAllUsers() (users []apiUser, err error) {
	cmd := `select id, name from users`
	rows, err := Db.Query(cmd)
	if err != nil {
		log.Fatalln(err)
	}
	for rows.Next() {
		var user apiUser
		err = rows.Scan(&user.ID, &user.Name)
		if err != nil {
			log.Fatalln(err)
		}
		users = append(users, user)
	}
	rows.Close()
	return users, err
}

func GetUser(id int) (user User, err error) {
	user = User{}
	cmd := `select id, name, password, created_at from users where id = ?`
	err = Db.QueryRow(cmd, id).Scan(
		&user.ID,
		&user.Name,
		&user.PassWord,
		&user.CreatedAt,
	)
	return user, err
}

func GetApiUser(name string) (user apiUser, err error) {
	user = apiUser{}
	cmd := `select id, name from users where name = ?`
	err = Db.QueryRow(cmd, name).Scan(
		&user.ID,
		&user.Name,
	)
	return user, err
}

func GetUserByUserName(user_name string) (user User, err error) {
	user = User{}
	cmd := `select id, name, password, created_at from users where name = ?`
	err = Db.QueryRow(cmd, user_name).Scan(
		&user.ID,
		&user.Name,
		&user.PassWord,
		&user.CreatedAt,
	)
	return user, err
}

func (u *User) CreateSession() (session Session, err error) {
	session = Session{}
	cmd1 := `insert into sessions (user_id, name) values (?, ?)`
	_, err = Db.Exec(cmd1, u.ID, u.Name)
	if err != nil {
		log.Println(err)
	}
	cmd2 := `select id, user_id, created_at from sessions where user_id = ?`
	err = Db.QueryRow(cmd2, u.ID).Scan(&session.ID, &session.UserID, &session.CreatedAt)
	return session, err
}

func (s *Session) CheckSession() (valid bool, err error) {
	cmd := `select id, name, user_id, created_at from sessions where name = ?`
	err = Db.QueryRow(cmd, s.UserName).Scan(&s.ID, &s.UserName, &s.UserID, &s.CreatedAt)
	if err != nil {
		valid = false
		return
	}
	if s.ID != 0 {
		valid = true
	}
	return valid, err
}

func (s *Session) DeleteSession() (err error) {
	cmd := `delete from sessions where email = ?`
	_, err = Db.Exec(cmd, s.ID)
	if err != nil {
		log.Println(err)
	}
	return err
}

func (s *Session) GetUserBySession() (user User, err error) {
	user = User{}
	cmd := `select id, name, email, created_at from users where id = ?`
	err = Db.QueryRow(cmd, s.UserID).Scan(&user.ID, &user.Name, &user.CreatedAt)
	return user, err
}
