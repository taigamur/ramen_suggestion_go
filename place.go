package main

import "log"

type Place struct {
	ID      int
	Name    string
	Address string
	Value   int
}

func CreatePlace(id int, place string, address string) (err error) {
	cmd := `insert into places(
		id,
		name,
		address
	) values (?, ?, ?) ON DUPLICATE KEY UPDATE name=VALUES(name), address=VALUES(address)`

	_, err = Db.Exec(cmd,
		id,
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

func create_init_place() {
	CreatePlace(1, "place1", "つくば市1")
	CreatePlace(2, "place1", "つくば市2")
	CreatePlace(3, "place3", "つくば市3")
	CreatePlace(4, "place4", "土浦市1")
	CreatePlace(5, "place5", "つくば市5")
}
