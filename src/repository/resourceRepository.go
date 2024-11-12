/*
 * Resource Inventory Management
 * API version: 4.0.0
 */

package repository

import (
	"context"
	"fmt"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"resource-inventory/db"
	"resource-inventory/src/models"
	"time"
)

const ResourceCollection = "Resources"

func getDB() *mongo.Database {
	return db.DBinstance()
}

func ListResources(filterMap bson.M, findOptions *options.FindOptions) ([]models.Resource, error) {
	log.Debug("ListResource repository started")

	database := getDB()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := database.Collection("Resources").Find(ctx, filterMap, findOptions)
	if err != nil {
		return nil, fmt.Errorf("error fetching resources from database: %v", err)
	}
	defer func(cursor *mongo.Cursor, ctx context.Context) {
		err := cursor.Close(ctx)
		if err != nil {

		}
	}(cursor, ctx)

	var resources []models.Resource
	if err = cursor.All(ctx, &resources); err != nil {
		return nil, fmt.Errorf("error reading resources from cursor: %v", err)
	}

	log.Debug("ListResource repository completed")
	return resources, nil
}

func CreateResource(resource models.Resource) (models.Resource, error) {
	log.Debug("CreateResource repository started")

	database := getDB()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	res, err := database.Collection(ResourceCollection).InsertOne(ctx, resource)
	if err != nil {
		return models.Resource{}, err
	}

	resource.Id = res.InsertedID.(string)

	log.Debug("CreateResource repository completed")
	return resource, nil
}

func RetrieveResource(id string) (models.Resource, error) {
	log.Debug("RetrieveResource repository started")

	database := getDB()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var resource models.Resource
	filter := bson.M{"id": id}
	err := database.Collection(ResourceCollection).FindOne(ctx, filter).Decode(&resource)
	if err != nil {
		return models.Resource{}, err
	}

	log.Debug("RetrieveResource repository completed")
	return resource, nil
}

func PatchResource(id string, update models.Resource) (models.Resource, error) {
	log.Debug("PatchResource repository started")

	database := getDB()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := database.Collection(ResourceCollection).UpdateOne(ctx, bson.M{"id": id}, bson.D{{"$set", update}})
	if err != nil {
		return models.Resource{}, err
	}

	updatedResource, err := RetrieveResource(id)
	if err != nil {
		return models.Resource{}, err
	}

	log.Debug("PatchResource repository completed")
	return updatedResource, nil
}

func DeleteResource(id string) error {
	log.Debug("DeleteResource repository started")

	database := getDB()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{"id": id}
	_, err := database.Collection(ResourceCollection).DeleteOne(ctx, filter)

	log.Debug("DeleteResource repository completed")
	return err
}
