package main

import (
	"log"

	"github.com/fetch-rewards/receipt-processor-challenge/pkg/handlers"
	"github.com/fetch-rewards/receipt-processor-challenge/pkg/repositories"
	"github.com/fetch-rewards/receipt-processor-challenge/pkg/services"
	"github.com/fetch-rewards/receipt-processor-challenge/pkg/validation"
	"github.com/gin-gonic/gin"
)

func createHandler() *handlers.ReceiptHandler {
	repo := repositories.NewInMemoryReceiptPointsRepository()
	svc := services.NewDefaultPointsService()
	return handlers.NewReceiptHandler(repo, svc)
}

func main() {
	engine := gin.Default()

	err := validation.SetupCustomValidationRules()
	if err != nil {
		log.Panicf("failed to register custom validation rules: %v", err)
	}

	router := engine.Group("receipts")
	h := createHandler()
	h.SetupRoutes(router)
	if err := engine.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
