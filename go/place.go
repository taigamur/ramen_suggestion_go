package main

import "log"

type Place struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Address string `json:"address"`
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

func GetPlace(id int) (place Place, err error) {
	place = Place{}
	cmd := `select id, name, palce from places where id = ?`
	err = Db.QueryRow(cmd, id).Scan(&place.ID, &place.Name, &place.Address)
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
	log.Println("s: " + s)
	rows, err := Db.Query(cmd, s)

	for rows.Next() {
		var place Place
		err = rows.Scan(&place.ID, &place.Name, &place.Address)
		places = append(places, place)
	}
	rows.Close()
	return places, err
}
