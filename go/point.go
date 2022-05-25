package main

import "log"

type Point struct {
	ID       int
	UserName string
	PlaceID  int
	Value    int
}

func GetPoint(username string, place_id int) (point Point) {
	cmd := `select id, username, place_id, value from points where username = ? and place_id = ?`
	err := Db.QueryRow(cmd, username, place_id).Scan(
		&point.ID,
		&point.UserName,
		&point.PlaceID,
		&point.Value,
	)
	if err != nil {
		log.Println(err)
	}
	return point
}

func CreatePoint(username string, place_id int, value int) (err error) {

	cmd := `insert into points(
		username,
		place_id,
		value
	) values (?, ?, ?) `
	_, err = Db.Exec(cmd, username, place_id, value)
	if err != nil {
		log.Println(err)
	}
	return err
}

func UpdatePoint(username string, place_id int, value int) (err error) {
	cmd := `update points set value = ? where username = ? and place_id = ?`
	_, err = Db.Exec(cmd, value, username, place_id)

	if err != nil {
		log.Println(err)
	}
	return err
}
