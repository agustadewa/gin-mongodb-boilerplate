package controller

import (
	"context"
	"fmt"
	"net/http"

	"github.com/agustadewa/gin-mongodb-boilerplate/service"
	"github.com/gin-gonic/gin"
)

// Controller is a controller object that have services
type Controller struct {
	// user service
	user *service.Service
}

// New generates new controller object
func New(ctx context.Context) (*Controller, error) {
	// create user service
	userSVC, err := service.New(ctx)
	if err != nil {
		return nil, fmt.Errorf("error creating service: %v", err)
	}

	return &Controller{
		user: userSVC,
	}, nil
}

// Check returns ok
func (ctrl *Controller) Check(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}
