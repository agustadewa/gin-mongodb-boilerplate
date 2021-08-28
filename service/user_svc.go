package service

import (
	"context"
	"fmt"

	"github.com/agustadewa/gin-mongodb-boilerplate/repository"
	"github.com/agustadewa/gin-mongodb-boilerplate/utils"
)

// Service is a service object that have repos
type Service struct {
	// user mongodb repository
	db *repository.MongoRepo
}

// New generates new service object
func New(ctx context.Context) (*Service, error) {
	// get mongo uri config
	mongoURI := ""

	// create mongodb connection
	mongoClient, err := utils.NewMongoConnection(ctx, mongoURI)
	if err != nil {
		return nil, fmt.Errorf("error connecting to mongodb: %v", err)
	}

	return &Service{
		db: repository.New(mongoClient),
	}, nil
}
