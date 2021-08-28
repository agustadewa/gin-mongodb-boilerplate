package main

import (
	"context"
	"fmt"

	"github.com/agustadewa/gin-mongodb-boilerplate/controller"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func main() {
	// create controller
	ctrl, err := controller.New(context.Background())
	if err != nil {
		logrus.Fatalf("error creating controller: %v", err)
	}

	r := gin.Default()

	// Check
	r.GET("/test", ctrl.Check)

	// get gin config
	host := ""
	port := ""

	if err = r.Run(fmt.Sprintf("%v:%v", host, port)); err != nil {
		logrus.Fatalf("error running gin: %v", err)
	}
}
