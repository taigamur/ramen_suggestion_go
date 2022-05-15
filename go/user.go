package main

import (
	"log"
	"time"
)

type User struct {
	ID        int
	Name      string
	Email     string
	PassWord  string
	CreatedAt time.Time
	Posts     []Post
	Places    []Place
}

type apiUser struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type Session struct {
	ID        int
	Email     string
	UserID    int
	CreatedAt time.Time
}

func (u *User) CreateUser() (err error) {
	cmd := `insert into users (
		name,
		email,
		password) values (?, ?, ?)`

	_, err = Db.Exec(cmd,
		u.Name,
		u.Email,
		Encrypt(u.PassWord))

	if err != nil {
		log.Println(err)
	}
	return err
}

// React 連携確認用のテスト
func GetAllUsers() (users []apiUser, err error) {
	cmd := `select id, name, email from users`
	rows, err := Db.Query(cmd)
	if err != nil {
		log.Fatalln(err)
	}
	for rows.Next() {
		var user apiUser
		err = rows.Scan(&user.ID, &user.Name, &user.Email)
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
	cmd := `select id, name, email, password, created_at from users where id = ?`
	err = Db.QueryRow(cmd, id).Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.PassWord,
		&user.CreatedAt,
	)
	return user, err
}

func GetUserByEmail(email string) (user User, err error) {
	user = User{}
	cmd := `select id, name, email, password, created_at from users where email = ?`
	err = Db.QueryRow(cmd, email).Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.PassWord,
		&user.CreatedAt,
	)
	return user, err
}

func (u *User) CreateSession() (session Session, err error) {
	session = Session{}
	cmd1 := `insert into sessions (email, user_id) values (?,?)`
	_, err = Db.Exec(cmd1, u.Email, u.ID)
	if err != nil {
		log.Println(err)
	}
	cmd2 := `select id, email, user_id, created_at from sessions where user_id = ? and email = ?`
	err = Db.QueryRow(cmd2, u.ID, u.Email).Scan(&session.ID, &session.Email, &session.UserID, &session.CreatedAt)
	return session, err
}

func (s *Session) CheckSession() (valid bool, err error) {
	cmd := `select id, email, user_id, created_at from sessions where email = ?`
	err = Db.QueryRow(cmd, s.Email).Scan(&s.ID, &s.Email, &s.UserID, &s.CreatedAt)
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
	err = Db.QueryRow(cmd, s.UserID).Scan(&user.ID, &user.Name, &user.Email, &user.CreatedAt)
	return user, err
}
