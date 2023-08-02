package models

type Movie struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Genre  string `json:"genre"`
	Rating int    `json:"rating"`
}
