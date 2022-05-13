package main

import (
	"crypto/sha1"
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"
)

var Db *sql.DB
var err error

const (
	tableNameUser    = "users"
	tableNamePost    = "posts"
	talbeNameSession = "sessions"
	tableNamePlace   = "places"
	tableNamePoint   = "points"
)

func init() {

	Db := connectDB()
	// defer Db.Close()

	cmdU := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s(
	id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
	name TEXT,
	email varchar(255) UNIQUE KEY,
	password TEXT,
	created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP)`, tableNameUser)
	Db.Exec(cmdU)

	cmdP := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s(
		id INT AUTO_INCREMENT PRIMARY KEY,
		user_id INT NOT NULL,
		place_id INT NOT NULL,
		value INT,
		comment TEXT,
		date DATETIME NOT NULL,
		created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP)`, tableNamePost)
	Db.Exec(cmdP)

	cmdS := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s(
		id INT AUTO_INCREMENT PRIMARY KEY,
		email TEXT,
		user_id INT,
		created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP)`, talbeNameSession)
	Db.Exec(cmdS)

	cmd := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s(
		id INT PRIMARY KEY,
		name TEXT,
		address TEXT)`, tableNamePlace)
	Db.Exec(cmd)
	create_init_place()

	cmdT := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s(
		id INT AUTO_INCREMENT PRIMARY KEY,
		user_id INT NOT NULL,
		place_id INT NOT NULL,
		value INT,
		count INT)`, tableNamePoint)
	Db.Exec(cmdT)
}

func open(path string, count uint) *sql.DB {
	Db, err = sql.Open("mysql", path)
	if err != nil {
		log.Fatal("open error:", err)
	}

	if err = Db.Ping(); err != nil {
		time.Sleep(time.Second * 2)
		count--
		fmt.Printf("retry... count:%v\n", count)
		return open(path, count)
	}

	fmt.Println("db connected!!")
	return Db
}

func connectDB() *sql.DB {
	var path string = fmt.Sprintf("%s:%s@tcp(db:3306)/%s?charset=utf8&parseTime=true",
		os.Getenv("MYSQL_USER"), os.Getenv("MYSQL_PASSWORD"),
		os.Getenv("MYSQL_DATABASE"))

	return open(path, 100)
}

func Encrypt(plaintext string) (cryptext string) {
	cryptext = fmt.Sprintf("%x", sha1.Sum([]byte(plaintext)))
	return cryptext
}