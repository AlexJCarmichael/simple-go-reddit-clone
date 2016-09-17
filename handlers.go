package main

import (
	"encoding/json"
 	"net/http"

  "github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/postgres"

	"github.com/gorilla/mux"
)

func LinkIndex(w http.ResponseWriter, r *http.Request) {
	links := Links{
		Link{Url: "www.google.com", Title: "My cool search engine"},
		Link{Url: "www.bing.com", Title: "My other cool search engine"},
	}
  w.Header().Set("Content-Type", "application/json; charset=UTF-8")
  w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(links); err != nil {
		panic(err)
	}
}

func LinkShow(w http.ResponseWriter, r *http.Request) {
  db, _ := gorm.Open("postgres", "host=localhost dbname=go-api-db")
	params := mux.Vars(r)
	linkId := params["linkId"]
  w.WriteHeader(http.StatusOK)
  var link Link
  show := db.First(&link, linkId)
	if err := json.NewEncoder(w).Encode(show); err != nil {
    panic(err)
  }
  defer db.Close()
}
