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
	CreatedAt time.Time
}

func (u *User) CreatePost(place_id int, value int, comment string) (err error) {
	cmd := `insert into posts (
		user_id,
		place_id,
		value,
		comment) values (?,?,?,?)`
	_, err = Db.Exec(cmd, u.ID, place_id, value, comment)
	if err != nil {
		log.Fatalln(err)
	}
	return err
}

func GetPost(id int) (post Post, err error) {
	cmd := `select id, user_id, place_id, value, comment, created_at from posts where id = ?`
	err = Db.QueryRow(cmd, id).Scan(
		&post.ID,
		&post.UserID,
		&post.PlaceID,
		&post.Value,
		&post.Comment,
		&post.CreatedAt,
	)
	return post, err
}

func GetPosts(user_id int) (posts []Post, err error) {
	cmd := `select id, user_id, place_id, value, comment, created_at from posts where user_id = ?`
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
	cmd := `update posts set comment = ?, place_id = ?, value = ? , user_id = ? where id = ?`
	_, err = Db.Exec(cmd, t.Comment, t.PlaceID, t.Value, t.UserID, t.ID)
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
