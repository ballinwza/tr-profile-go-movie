package handler_movie

import (
	connecter_mongodb "github.com/ballinwza/tr-profile-go-movie/connecter"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type MovieService struct {
	collection *mongo.Collection
}

func SetupMovieService() *MovieService {

	return &MovieService{
		collection: connecter_mongodb.ConnectWithMongoDb("tr-profile-go-movie", "my-favorite-movies"),
	}
}
