package db

import (
	"context"
	"fmt"
	log "github.com/sirupsen/logrus"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// DBinstance func
func DBinstance() *mongo.Database {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	MongoDb := os.Getenv("MONGODB_URL")
	DbName := os.Getenv("DB_NAME")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(MongoDb))
	if err != nil {
		log.Fatal("Failed to connect to MongoDB:", err)
	}

	fmt.Println("Connected to MongoDB!")
	db := client.Database(DbName)
	return db
}
