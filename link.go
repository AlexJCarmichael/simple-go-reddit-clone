package main

import "time"

type Link struct {
	Url     string    `json:"url"`
	Title   string    `json:"title"`
	Created time.Time `json:"created"`
}

type Links []Link
