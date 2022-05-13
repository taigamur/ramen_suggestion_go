package main

import (
	"log"
	"time"
)

type Post struct {
	ID        int
	UserID    int
	PlaceID   int
	Value     int
	Comment   string
	Date      time.Time
	CreatedAt time.Time
}

func (u *User) CreatePost(place_id int, value int, comment string, date time.Time) (err error) {
	cmd := `insert into posts (
		user_id,
		place_id,
		value,
		comment,
		date) values (?,?,?,?,?)`
	_, err = Db.Exec(cmd, u.ID, place_id, value, comment, date)
	if err != nil {
		log.Println(err)
	}
	return err
}

func GetPost(id int) (post Post, err error) {
	cmd := `select id, user_id, place_id, value, comment, date, created_at from posts where id = ?`
	err = Db.QueryRow(cmd, id).Scan(
		&post.ID,
		&post.UserID,
		&post.PlaceID,
		&post.Value,
		&post.Comment,
		&post.Date,
		&post.CreatedAt,
	)
	return post, err
}

func GetPosts(user_id int) (posts []Post, err error) {
	cmd := `select id, user_id, place_id, value, comment, date, created_at from posts where user_id = ? order by date desc`
	rows, err := Db.Query(cmd, user_id)
	if err != nil {
		log.Fatalln(err)
	}
	for rows.Next() {
		var post Post
		err = rows.Scan(&post.ID,
			&post.UserID,
			&post.PlaceID,
			&post.Value,
			&post.Comment,
			&post.Date,
			&post.CreatedAt,
		)
		if err != nil {
			log.Fatalln(err)
		}
		posts = append(posts, post)
	}
	rows.Close()
	return posts, err
}

func (t *Post) UpdatePost() error {
	cmd := `update posts set comment = ?, place_id = ?, value = ? , user_id = ? , date = ? where id = ?`
	_, err = Db.Exec(cmd, t.Comment, t.PlaceID, t.Value, t.UserID, t.Date, t.ID)
	if err != nil {
		log.Fatalln(err)
	}
	return err
}

func (t *Post) DeletePost() error {
	cmd := `delete from posts where id = ?`
	_, err = Db.Exec(cmd, t.ID)
	if err != nil {
		log.Fatalln(err)
	}
	return err
}

func (p *Post) GetMonth() int {
	return int(p.Date.Month())
}

func (p *Post) GetDay() int {
	return int(p.Date.Day())
}
