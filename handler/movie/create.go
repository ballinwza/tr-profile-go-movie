package handler_movie

import (
	"context"
	"time"

	struct_movie "github.com/ballinwza/tr-profile-go-movie/struct"
	"github.com/gofiber/fiber/v2"
)

func (m *MovieService) CreateMovieHandler(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	movie := struct_movie.Movie{}
	err := c.BodyParser(&movie)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Error CreateMovieHandler getBody value",
			"error":   err.Error(),
		})
	}

	res, err := m.collection.InsertOne(ctx, movie)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Error CreateMovieHandler InsertOne.",
			"error":   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"id": res.InsertedID,
	})
}
