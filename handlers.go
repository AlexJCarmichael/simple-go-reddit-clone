package main

import (
	"encoding/json"
	"fmt"
	"net/http"

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
	params := mux.Vars(r)
	linkId := params["linkId"]
  w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Link show:", linkId)
}
