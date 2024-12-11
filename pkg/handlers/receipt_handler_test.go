package handlers

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/fetch-rewards/receipt-processor-challenge/pkg/models"
	"github.com/fetch-rewards/receipt-processor-challenge/pkg/repositories"
	"github.com/fetch-rewards/receipt-processor-challenge/pkg/services"
	"github.com/fetch-rewards/receipt-processor-challenge/pkg/validation"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func callEngine(t *testing.T, engine *gin.Engine, method string, path string, data any, expectedStatus int, body any) {
	var bodyData io.Reader
	if data != nil {
		buf, err := json.Marshal(data)
		assert.NoError(t, err)
		bodyData = bytes.NewBuffer(buf)
	}

	req, err := http.NewRequest(method, path, bodyData)
	assert.NoError(t, err)

	w := httptest.NewRecorder()
	engine.ServeHTTP(w, req)

	assert.Equal(t, expectedStatus, w.Code)

	if body != nil {
		err = json.Unmarshal(w.Body.Bytes(), body)
		assert.NoError(t, err)
	}
}

func TestProcessEndpoint(t *testing.T) {
	repo := repositories.NewInMemoryReceiptPointsRepository()
	svc := services.NewDefaultPointsService()
	h := NewReceiptHandler(repo, svc)

	gin.SetMode(gin.TestMode)
	engine := gin.Default()
	err := validation.SetupCustomValidationRules()
	assert.NoError(t, err)

	router := engine.Group("receipts")
	h.SetupRoutes(router)

	var body gin.H
	callEngine(t, engine, "POST", "/receipts/process", &models.Receipt{
		Retailer:     "Target",
		PurchaseDate: "2022-01-02",
		PurchaseTime: "13:13",
		Total:        "1.25",
		Items: []models.ReceiptItem{
			{
				ShortDescription: "Pepsi - 12-oz",
				Price:            "1.25",
			},
		},
	}, http.StatusOK, &body)
	assert.Contains(t, body, "id")

	id, ok := body["id"].(string)
	assert.True(t, ok)

	_, err = repo.GetReceiptPointsById(id)
	assert.NoError(t, err)
}

func TestGetPoints(t *testing.T) {
	repo := repositories.NewInMemoryReceiptPointsRepository()
	svc := services.NewDefaultPointsService()
	h := NewReceiptHandler(repo, svc)

	gin.SetMode(gin.TestMode)
	engine := gin.Default()
	err := validation.SetupCustomValidationRules()
	assert.NoError(t, err)

	router := engine.Group("receipts")
	h.SetupRoutes(router)

	// insert one item in repo
	id, err := repo.SaveReceiptPoints(&models.Receipt{Retailer: "Test"}, 100)
	assert.NoError(t, err)

	var body gin.H
	callEngine(t, engine, "GET", "/receipts/"+id+"/points", nil, http.StatusOK, &body)
	assert.Contains(t, body, "points")
	assert.Equal(t, 100.0, body["points"])
}

func TestInvalidId(t *testing.T) {
	repo := repositories.NewInMemoryReceiptPointsRepository()
	svc := services.NewDefaultPointsService()
	h := NewReceiptHandler(repo, svc)

	gin.SetMode(gin.TestMode)
	engine := gin.Default()
	err := validation.SetupCustomValidationRules()
	assert.NoError(t, err)

	router := engine.Group("receipts")
	h.SetupRoutes(router)

	callEngine(t, engine, "GET", "/receipts/some-id/points", nil, http.StatusNotFound, nil)
}
