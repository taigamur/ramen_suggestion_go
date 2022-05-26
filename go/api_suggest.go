package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"
)

func suggestPlace(w http.ResponseWriter, r *http.Request) {
	setApiHeader(w)

	// suggest リストの取得
	username := r.URL.Query().Get("username")
	suggests, _ := GetSuggestPlaces(username)
	log.Println(suggests)

	// 総和の計算
	sum := 0
	for _, s := range suggests {
		sum += s.Value
	}
	// 乱数の取得
	rand.Seed(time.Now().UnixNano())
	rand_n := rand.Intn(sum)

	// 対応するplaceの取得
	val_sum := 0
	var suggest = Suggest{}
	for _, s := range suggests {
		val_sum += s.Value
		if val_sum >= rand_n {
			suggest = s
			break
		}
	}

	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		log.Println(err)
	} else {
		w.WriteHeader(http.StatusOK)
		res, _ := json.Marshal(suggest)
		fmt.Fprint(w, string(res))
	}
}
