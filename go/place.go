package main

import (
	"database/sql"
	"log"
)

type Place struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Address string `json:"address"`
}

type Suggest struct {
	Place Place `json:"place"`
	Value int   `json:"value"`
	Flag  bool  `json:"flag"`
}

func CreatePlace(id int, place string, hiragana string, address string) (err error) {
	cmd := `insert into places(
		id,
		name,
		hiragana,
		address
	) values (?, ?, ?, ?) ON DUPLICATE KEY UPDATE name=VALUES(name), hiragana=VALUES(hiragana), address=VALUES(address)`

	_, err = Db.Exec(cmd,
		id,
		place,
		hiragana,
		address)
	if err != nil {
		log.Fatalln(err)
	}
	return err
}

func GetPlace(id int) (place Suggest, err error) {
	place = Suggest{}
	cmd := `select id, name, address from places where id = ?`
	err = Db.QueryRow(cmd, id).Scan(&place.Place.ID, &place.Place.Name, &place.Place.Address)
	place.Value = 5
	place.Flag = false
	return place, err
}

func GetPlaces() (places []Place, err error) {
	cmd := `select id, name, address from places`
	rows, err := Db.Query(cmd)
	if err != nil {
		log.Fatalln(err)
	}
	for rows.Next() {
		var place Place
		err = rows.Scan(&place.ID,
			&place.Name,
			&place.Address,
		)
		if err != nil {
			log.Fatalln(err)
		}
		places = append(places, place)
	}
	rows.Close()
	return places, err
}

func GetPlacesByKeyword(keyword string) (places []Place, err error) {
	cmd := `select id, name, address from places where concat(name, hiragana) like ?`
	s := "%" + keyword + "%"
	rows, err := Db.Query(cmd, s)

	for rows.Next() {
		var place Place
		err = rows.Scan(&place.ID, &place.Name, &place.Address)
		places = append(places, place)
	}
	rows.Close()
	return places, err
}

func GetSuggestPlaces(username string) (suggests []Suggest, err error) {
	cmd := `select id, name, address from places`
	rows, err := Db.Query(cmd)
	if err != nil {
		log.Println(err)
	}
	for rows.Next() {
		var suggest Suggest
		err = rows.Scan(
			&suggest.Place.ID,
			&suggest.Place.Name,
			&suggest.Place.Address,
		)
		if err != nil {
			log.Println(err)
		}
		cmdCnt := `select value from points where place_id = ? AND username = ? `
		err := Db.QueryRow(cmdCnt, suggest.Place.ID, username).Scan(
			&suggest.Value,
		)
		if err == sql.ErrNoRows {
			suggest.Flag = false
			suggest.Value = 5
		} else if err != nil {
			log.Println(err)
		} else {
			suggest.Flag = true
		}

		suggests = append(suggests, suggest)
	}
	rows.Close()
	return suggests, err
}
