
package main

import (
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/postgres"
)

func db() {
  db, err := gorm.Open("sqlite3", "/tmp/gorm.db")
  if err != nil {
    panic("Failed to connect database")
  }
  db.AutoMigrate(&Link{})
  defer db.Close()
}
