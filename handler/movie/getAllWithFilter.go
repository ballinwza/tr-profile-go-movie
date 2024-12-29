package handler_movie

import (
	"context"
	"strconv"

	struct_movie "github.com/ballinwza/tr-profile-go-movie/struct"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func (m *MovieService) GetAllMovieWithFilterHandler(c *fiber.Ctx) error {
	var result []struct_movie.Movie
	qAdult := c.Query("adult")

	opts := options.Find().SetSort(bson.M{"title": 1})

	filter := bson.M{}

	if qAdult != "" {
		boolQAdult, err := strconv.ParseBool(qAdult)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Error at GetAllMovieWithFilterHandler convert string to bool.",
				"error":   err.Error(),
			})
		}
		filter = bson.M{
			"adult": boolQAdult,
		}
	}

	cursor, err := m.collection.Find(context.TODO(), filter, opts)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error at GetAllMovieWithFilterHandler Find.",
			"error":   err.Error(),
		})
	}

	if err = cursor.All(context.TODO(), &result); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error at GetAllMovieWithFilterHandler fill list with movie.",
			"error":   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(result)
}
