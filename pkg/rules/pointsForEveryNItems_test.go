package rules

import (
	"testing"

	"github.com/fetch-rewards/receipt-processor-challenge/pkg/models"
	"github.com/stretchr/testify/assert"
)

func TestPointsForEveryNItemsNotEnoughItems(t *testing.T) {
	rule := NewDefaultPointsForEveryNItems()
	receipt := models.Receipt{
		Items: []models.ReceiptItem{
			{ShortDescription: "Pepsi - 12-oz", Price: "1.25"},
		},
	}

	desc := rule.DescribePoints(&receipt, true)
	assert.Equal(t, 0, desc.Points, "%v", desc.GetDescription())
}
func TestPointsForEveryNItems(t *testing.T) {
	rule := NewDefaultPointsForEveryNItems()
	receipt := models.Receipt{
		Items: []models.ReceiptItem{
			{ShortDescription: "Mountain Dew 12PK", Price: "6.49"},
			{ShortDescription: "Emils Cheese Pizza", Price: "15.25"},
			{ShortDescription: "Knorr Creamy Chicken", Price: "1.26"},
			{ShortDescription: "Doritos Nacho Cheese", Price: "3.35"},
			{ShortDescription: "   Klarbrunn 12-PK 12 FL OZ  ", Price: "12.00"},
		},
	}

	desc := rule.DescribePoints(&receipt, true)
	expected := (len(receipt.Items) / Default_PointsForEveryNItems_N) * Default_PointsForEveryNItems_Each
	assert.Equal(t, expected, desc.Points, "%v", desc.GetDescription())
}
