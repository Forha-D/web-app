package models

type Movie struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Genre  string `json:"genre"`
	Rating int    `json:"rating"`
}

type User struct {
	Name        string `json:"name,omitempty"  validate:"required"`
	Email       string `json:"email,omitempty"  validate:"required,email"`
	MovieDetail *Movie
}
