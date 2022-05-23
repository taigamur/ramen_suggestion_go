package main

import (
	"log"
	"time"
)

type Post struct {
	ID        int       `json:"id"`
	UserName  string    `json:"username"`
	PlaceID   int       `json:"place_id"`
	Value     int       `json:"value"`
	Date      string    `json:"date"`
	CreatedAt time.Time `json:"created_at"`
	Place     Place     `json:"place"`
}

func CreatePost(place_id int, username string, value int, date string) (err error) {
	cmd := `insert into posts (
		username,
		place_id,
		value,
		date) values (?,?,?,?)`
	_, err = Db.Exec(cmd, username, place_id, value, date)
	if err != nil {
		log.Println(err)
	}
	return err
}

func GetPost(id int) (post Post, err error) {
	cmd := `select id, username, place_id, value, date, created_at from posts where id = ?`
	err = Db.QueryRow(cmd, id).Scan(
		&post.ID,
		&post.UserName,
		&post.PlaceID,
		&post.Value,
		&post.Date,
		&post.CreatedAt,
	)
	return post, err
}

func GetPosts(username string) (posts []Post, err error) {
	cmd := `select id, username, place_id, value, date, created_at from posts where username = ? order by date desc`
	rows, err := Db.Query(cmd, username)
	if err != nil {
		log.Fatalln(err)
	}
	for rows.Next() {
		var post Post
		err = rows.Scan(
			&post.ID,
			&post.UserName,
			&post.PlaceID,
			&post.Value,
			&post.Date,
			&post.CreatedAt,
		)
		if err != nil {
			log.Fatalln(err)
		}

		cmdPlace := `select id, name, address from places where id = ?`
		err = Db.QueryRow(cmdPlace, post.PlaceID).Scan(
			&post.Place.ID,
			&post.Place.Name,
			&post.Place.Address,
		)
		if err != nil {
			log.Println(err)
		}

		posts = append(posts, post)
	}
	rows.Close()
	return posts, err
}

func (t *Post) UpdatePost() error {
	cmd := `update posts set place_id = ?, value = ? , username = ? , date = ? where id = ?`
	_, err = Db.Exec(cmd, t.PlaceID, t.Value, t.UserName, t.Date, t.ID)
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

// func (p *Post) GetMonth() int {
// 	return int(p.Date.Month())
// }

// func (p *Post) GetDay() int {
// 	return int(p.Date.Day())
// }
