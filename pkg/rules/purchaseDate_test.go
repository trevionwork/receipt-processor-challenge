package rules

import (
	"testing"

	"github.com/fetch-rewards/receipt-processor-challenge/pkg/models"
	"github.com/stretchr/testify/assert"
)

func TestPurchaseDate(t *testing.T) {
	rule := NewDefaultPurchaseDate()
	receipt := models.Receipt{
		PurchaseDate: "2022-01-01",
	}

	desc := rule.DescribePoints(&receipt, true)
	assert.Equal(t, Default_PurchaseDate_Points, desc.Points, "%v", desc.GetDescription())
}
func TestPurchaseDateEven(t *testing.T) {
	rule := NewDefaultPurchaseDate()
	receipt := models.Receipt{
		PurchaseDate: "2022-01-02",
	}

	desc := rule.DescribePoints(&receipt, true)
	assert.Equal(t, 0, desc.Points, "%v", desc.GetDescription())
}
