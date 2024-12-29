package connecter_mongodb

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func ConnectWithMongoDb(databaseName, collectionName string) *mongo.Collection {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	mongoUsername := os.Getenv("MONGO_USERNAME")
	mongoPassword := os.Getenv("MONGO_PASSWORD")

	conn := "mongodb+srv://" + mongoUsername + ":" + mongoPassword + "@user-cluster.xgu8qno.mongodb.net/?retryWrites=true&w=majority&appName=user-cluster"

	client, err := mongo.Connect(options.Client().ApplyURI(conn))
	if err != nil {
		log.Fatal(err)
		return nil
	}

	collection := client.Database(databaseName).Collection(collectionName)

	return collection

}
