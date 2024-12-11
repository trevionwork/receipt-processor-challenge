package handlers

import (
	"errors"
	"net/http"

	"github.com/fetch-rewards/receipt-processor-challenge/pkg/models"
	"github.com/fetch-rewards/receipt-processor-challenge/pkg/repositories"
	"github.com/fetch-rewards/receipt-processor-challenge/pkg/services"
	"github.com/gin-gonic/gin"
)

type ReceiptHandler struct {
	repo          repositories.ReceiptPointsRepository
	pointsService *services.PointsService
}

func NewReceiptHandler(repo repositories.ReceiptPointsRepository, pointsService *services.PointsService) *ReceiptHandler {
	return &ReceiptHandler{repo: repo, pointsService: pointsService}
}

func (h *ReceiptHandler) SetupRoutes(router *gin.RouterGroup) {
	router.POST("process", h.processReceipt)
	router.GET(":id/points", h.getPoints)
}
func (h *ReceiptHandler) processReceipt(c *gin.Context) {
	var receipt models.Receipt
	if err := c.ShouldBindBodyWithJSON(&receipt); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	data := gin.H{}
	var points int
	switch c.Query("desc") {
	case "full":
		desc := h.pointsService.DescribePoints(&receipt, true)
		data["desc"] = desc
		points = desc.TotalPoints

	case "compact":
		desc := h.pointsService.DescribePoints(&receipt, false)
		data["desc"] = desc
		points = desc.TotalPoints

	default:
		points = h.pointsService.ComputePoints(&receipt)
	}

	id, err := h.repo.SaveReceiptPoints(&receipt, points)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	data["id"] = id
	c.JSON(http.StatusOK, data)
}
func (h *ReceiptHandler) getPoints(c *gin.Context) {
	id := c.Param("id")

	points, err := h.repo.GetReceiptPointsById(id)
	if err != nil {
		if errors.Is(err, repositories.ErrNotFound) {
			c.Status(http.StatusNotFound)
			return
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"points": points})
}
