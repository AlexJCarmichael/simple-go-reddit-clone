package main

import "time"

type Link struct {
  ID        uint      `gorm:"primary_key"  json:"id"`
	Url       string    `json:"url"`
	Title     string    `json:"title"`
	CreatedAt time.Time `json:"created_at"`
  UpdatedAt time.Time `json:"updated_at"`
  DeletedAt time.Time `json:"deleted_at"`
}

type Links []Link
