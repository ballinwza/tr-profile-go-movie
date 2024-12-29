package main

import (
	"fmt"
	"log"

	handler_movie "github.com/ballinwza/tr-profile-go-movie/handler/movie"
	"github.com/gofiber/fiber/v2"
)

func main() {
	port := "3000"
	app := fiber.New()

	app.Get("/movies", handler_movie.SetupMovieService().GetAllMovieWithFilterHandler)
	app.Get("movie/:id", handler_movie.SetupMovieService().GetMovieByIdHandler)
	app.Post("/create/movie", handler_movie.SetupMovieService().CreateMovieHandler)

	fmt.Println("Server is running on port : ", port)
	log.Fatal(app.Listen(":" + port))

}
