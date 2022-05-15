package main

import "net/http"

// reactからのapiでcors関連のエラーに対処 headerにaccess-control-allow-originを与える
func setApiHeader(w http.ResponseWriter) error {
	// protocol := "http://"
	// host := "localhost:3000"
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, DELETE, PUT")
	w.Header().Set("Content-Type", "application/json")
	return nil
}
