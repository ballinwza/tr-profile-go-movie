package handler_movie

import (
	"context"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func (m *MovieService) DeleteMovieByIdHandler(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Error{
			Message: "Error at DeleteMovieByIdHandler parameter must be number",
			Code:    fiber.StatusBadRequest,
		})
	}

	res, err := m.collection.DeleteOne(context.TODO(), bson.M{"tmbd_id": id})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Error{
			Message: "Error at DeleteMovieByIdHandler DeleteOne",
			Code:    fiber.StatusInternalServerError,
		})
	}

	conId := strconv.Itoa(id)
	if res.DeletedCount == 0 {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Error{
			Message: "Error at DeleteMovieByIdHandler DeleteOne not founded movie Id : " + conId,
			Code:    fiber.StatusInternalServerError,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Movie Id : " + conId + " was deleted"})
}
