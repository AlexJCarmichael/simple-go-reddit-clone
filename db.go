
package main

import (
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/postgres"
)

func db() {
  db, err := gorm.Open("postgres", "host=localhost dbname=go-api-db")
  if err != nil {
    panic("Failed to connect database")
  }
  defer db.Close()
  db.CreateTable(&Link{})
  db.Create(&Link{Url: "www.google.com", Title: "My Cool Search Engine"})
}
