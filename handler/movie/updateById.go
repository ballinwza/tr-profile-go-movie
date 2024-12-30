package handler_movie

import (
	"context"
	"encoding/json"
	"strconv"

	struct_movie "github.com/ballinwza/tr-profile-go-movie/struct"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func (m *MovieService) UpdateMovieByIdHandler(c *fiber.Ctx) error {
	var bodyValue struct_movie.Movie

	body := c.Body()
	err := json.Unmarshal(body, &bodyValue)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Error{
			Code:    fiber.StatusBadRequest,
			Message: "Error UpdateMovieByIdHandler body value : " + err.Error(),
		})
	}

	var conIdToString string
	var filter bson.M
	update := bson.M{"$set": bodyValue}

	if bodyValue.TmdbId != nil {
		filter = bson.M{"tmbd_id": bodyValue.TmdbId}
		conIdToString = strconv.Itoa(*bodyValue.TmdbId)
	} else {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Error{
			Code:    fiber.StatusBadRequest,
			Message: "Error UpdateMovieByIdHandler not matched any movie tmbd_id : " + conIdToString,
		})
	}

	result, err := m.collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Error{
			Code:    fiber.StatusBadRequest,
			Message: "Error UpdateMovieByIdHandler updateOne : " + err.Error(),
		})
	}

	if result.MatchedCount == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Error{
			Code:    fiber.StatusBadRequest,
			Message: "Error UpdateMovieByIdHandler not matched any movie tmbd_id : " + conIdToString,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Matched and replace an existing movie",
		"data":    bodyValue,
	})

}
