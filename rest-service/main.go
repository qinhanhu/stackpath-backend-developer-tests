package main

import "C"
import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/stackpath/backend-developer-tests/rest-service/routers"
	"log"
)

func main() {
	fmt.Println("SP// Backend Developer Test - RESTful Service")
	fmt.Println()

	// TODO: Add RESTful web service here
	app := gin.Default()
	router := app.Group("")
	routers.AddRouters(router)

	err := app.Run(":8080")
	if err != nil {
		log.Fatalf("gin run error: %v", err)
		return
	}

}
