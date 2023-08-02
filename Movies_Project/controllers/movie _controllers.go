package controllers

import (
	"Movies_Project/models"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

var moviesDB = map[int]models.Movie{}

// Counter for generating movie IDs
var movieIDCounter = 1

func GetMovies(c echo.Context) error {
	var movies []models.Movie
	for _, movie := range moviesDB {
		movies = append(movies, movie)
	}
	return c.JSON(http.StatusOK, movies)
}

// Get a specific movie by ID
func GetMovie(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid movie ID")
	}

	movie, ok := moviesDB[id]
	if !ok {
		return echo.NewHTTPError(http.StatusNotFound, "Movie not found")
	}

	return c.JSON(http.StatusOK, movie)
}

// Create a new movie
func CreateMovie(c echo.Context) error {
	movie := new(models.Movie)
	if err := c.Bind(movie); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid data")
	}

	//fmt.Println(movie)

	movie.ID = movieIDCounter
	movieIDCounter++
	moviesDB[movie.ID] = *movie

	return c.JSON(http.StatusCreated, movie)
}

// Update an existing movie
func UpdateMovie(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid movie ID")
	}

	movie, ok := moviesDB[id]
	if !ok {
		return echo.NewHTTPError(http.StatusNotFound, "Movie not found")
	}

	updatedMovie := new(models.Movie)
	if err := c.Bind(updatedMovie); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid data")
	}

	movie.Title = updatedMovie.Title
	movie.Genre = updatedMovie.Genre
	movie.Rating = updatedMovie.Rating

	return c.JSON(http.StatusOK, movie)
}

// Delete a movie by ID
func DeleteMovie(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid movie ID")
	}

	_, ok := moviesDB[id]
	if !ok {
		return echo.NewHTTPError(http.StatusNotFound, "Movie not found")
	}

	delete(moviesDB, id)

	return c.NoContent(http.StatusNoContent)
}
