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
	tableNameUser = "users"
	tableNamePost = "posts"
)

/*
func init() {

	Db := connectDB()
	// defer Db.Close()

	cmdU := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s(
	id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
	name TEXT,
	email TEXT,
	password TEXT,
	created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP)`, tableNameUser)
	Db.Exec(cmdU)

	cmdP := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s(
		id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
		user_id INT NOT NULL,
		place_id INT NOT NULL,
		value INT,
		comment TEXT,
		created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP)`, tableNamePost)
	Db.Exec(cmdP)
}
*/

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
