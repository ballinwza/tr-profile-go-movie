package handler_movie

import (
	"context"
	"strconv"
	"time"

	struct_movie "github.com/ballinwza/tr-profile-go-movie/struct"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

func (m *MovieService) GetMovieByIdHandler(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	var result struct_movie.Movie
	pId := c.Params("id")
	id, err := strconv.Atoi(pId)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "error at GetMovieByIdHandler coverting string to int",
			"error":   err.Error(),
		})
	}

	err = m.collection.FindOne(ctx, bson.M{"tmbd_id": id}).Decode(&result)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "error at GetMovieByIdHandler findOne",
			"error":   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(result)
}
