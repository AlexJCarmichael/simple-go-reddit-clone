
package main

import (
  "fmt"
  "log"
  "net/http"

  "github.com/gorilla/mux"
)

func main() {
  router := mux.NewRouter().StrictSlash(true)
  router.HandleFunc("/", Root)
  router.HandleFunc("/links", LinkIndex)
  router.HandleFunc("/links/{linkId}", LinkShow)
  log.Fatal(http.ListenAndServe(":8080", router))
}


func Root(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "Landing page for my slim reddit")
}

func LinkIndex(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "Link Index Page")
}

func LinkShow(w http.ResponseWriter, r *http.Request) {
  params := mux.Vars(r)
  linkId := params["linkId"]
  fmt.Fprintf(w, "Link show:", linkId)
}
