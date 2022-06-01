package main

import (
	"database/sql"
	"log"
	"sort"
)

type Point struct {
	ID       int
	UserName string
	PlaceID  int
	Value    int
}

type PointApi struct {
	ID    int   `json:"id"`
	Place Place `json:"place"`
	Value int   `json:"value"`
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

// MyPointsに表示するためのデータ
func GetPoints(username string) (points []PointApi) {
	cmd := `select id, name, address from places`
	rows, err := Db.Query(cmd)

	if err != nil {
		log.Println("get points error")
	}
	// userのpoint一覧を取得
	for rows.Next() {
		var point PointApi
		err = rows.Scan(
			&point.Place.ID,
			&point.Place.Name,
			&point.Place.Address,
		)
		if err != nil {
			log.Println("get point err")
		}

		cmdPoint := `select id, value from points where username = ? and place_id = ?`
		err = Db.QueryRow(cmdPoint, username, point.Place.ID).Scan(&point.ID, &point.Value)
		if err == sql.ErrNoRows {
			point.Value = -1
			point.ID = -1
		} else if err != nil {
			log.Println(err)
		}
		points = append(points, point)
	}
	rows.Close()
	// valueの降順でsort
	sort.Slice(points, func(i, j int) bool { return points[i].Value > points[j].Value })

	return points
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
