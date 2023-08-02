package main

import (
	"Movies_Project/controllers"
	"github.com/labstack/echo/v4"
)

func main() {

	e := echo.New()
	e.GET("/movies", controllers.GetMovies)
	e.GET("/movies/:id", controllers.GetMovie)
	e.POST("/movies", controllers.CreateMovie)
	e.PUT("/movies/:id", controllers.UpdateMovie)
	e.DELETE("/movies/:id", controllers.DeleteMovie)

	e.Logger.Fatal(e.Start(":8000"))

}
