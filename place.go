package main

import "log"

type Place struct {
	ID      int
	Name    string
	Address string
	Value   int
}

func CreatePlace(place string, address string) (err error) {
	cmd := `insert into places(
		name,
		address
	) values (?, ?)`

	_, err = Db.Exec(cmd,
		place,
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
