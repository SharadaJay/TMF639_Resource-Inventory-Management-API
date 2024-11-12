/*
 * Resource Inventory Management
 * API version: 4.0.0
 */

package services

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"resource-inventory/src/models"
	"resource-inventory/src/repository"
	"strconv"
	"strings"
)

func ListResources(fields, limit, offset, filterByValString string) ([]models.Resource, error) {
	log.Debug("ListResource service started")

	filterMap, err := prepareFilterMap(filterByValString)
	if err != nil {
		return nil, err
	}

	findOptions, err := prepareFindOptions(fields, limit, offset)
	if err != nil {
		return nil, err
	}

	resources, err := repository.ListResources(filterMap, findOptions)
	if err != nil {
		return nil, fmt.Errorf("error fetching resources: %v", err)
	}

	log.Debug("ListResource service completed")
	return resources, nil
}

func CreateResource(resource models.Resource) (models.Resource, error) {
	return repository.CreateResource(resource)
}

func RetrieveResource(id string) (models.Resource, error) {
	return repository.RetrieveResource(id)
}

func PatchResource(id string, update models.Resource) (models.Resource, error) {
	return repository.PatchResource(id, update)
}

func DeleteResource(id string) error {
	return repository.DeleteResource(id)
}

func prepareFilterMap(filterByValString string) (bson.M, error) {
	var filterMap bson.M
	if len(filterByValString) != 0 {
		queryMap := make(map[string]string)
		filterByValList := strings.Split(filterByValString, "&")
		for _, v := range filterByValList {
			tempList := strings.Split(v, "=")
			queryMap[tempList[0]] = tempList[1]
		}

		data, err := bson.Marshal(queryMap)
		if err != nil {
			return nil, fmt.Errorf("error creating filter map: %v", err)
		}

		err = bson.Unmarshal(data, &filterMap)
		if err != nil {
			return nil, fmt.Errorf("error unmarshalling filter map: %v", err)
		}
	}
	return filterMap, nil
}

func prepareFindOptions(fields, limit, offset string) (*options.FindOptions, error) {
	findOptions := options.Find()

	if len(fields) != 0 {
		fieldList := strings.Split(fields, ",")
		fieldMap := make(map[string]int)
		for _, i := range fieldList {
			fieldMap[i] = 1
		}
		data, err := bson.Marshal(fieldMap)
		if err != nil {
			return nil, fmt.Errorf("error creating projection options: %v", err)
		}

		var optionsMap bson.M
		err = bson.Unmarshal(data, &optionsMap)
		if err != nil {
			return nil, fmt.Errorf("error unmarshalling projection options: %v", err)
		}

		findOptions.SetProjection(optionsMap)
	}

	if len(limit) != 0 {
		intLimit, err := strconv.Atoi(limit)
		if err != nil {
			return nil, fmt.Errorf("invalid limit: %v", err)
		}
		findOptions.SetLimit(int64(intLimit))
	}

	if len(offset) != 0 {
		intOffset, err := strconv.Atoi(offset)
		if err != nil {
			return nil, fmt.Errorf("invalid offset: %v", err)
		}
		findOptions.SetSkip(int64(intOffset))
	}

	return findOptions, nil
}
