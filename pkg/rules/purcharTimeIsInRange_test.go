package rules

import (
	"strings"
	"testing"

	"github.com/fetch-rewards/receipt-processor-challenge/pkg/models"
	"github.com/stretchr/testify/assert"
)

func TestNewPurchaseTimeIsInRangeBeforeTime(t *testing.T) {
	rule := NewDefaultPurchaseTimeIsInRange()
	receipt := models.Receipt{
		PurchaseTime: "13:55",
	}

	desc := rule.DescribePoints(&receipt, true)
	assert.Equal(t, 0, desc.Points, "%v", desc.GetDescription())
	assert.Equal(t, 1, len(desc.Description))
	assert.True(t, strings.Contains(desc.Description[0], "is before"))
}
func TestNewPurchaseTimeIsInRangeAfterTime(t *testing.T) {
	rule := NewDefaultPurchaseTimeIsInRange()
	receipt := models.Receipt{
		PurchaseTime: "16:05",
	}

	desc := rule.DescribePoints(&receipt, true)
	assert.Equal(t, 0, desc.Points, "%v", desc.GetDescription())
	assert.Equal(t, 1, len(desc.Description))
	assert.True(t, strings.Contains(desc.Description[0], "is after"))
}
func TestNewPurchaseTimeIsInRange(t *testing.T) {
	rule := NewDefaultPurchaseTimeIsInRange()
	receipt := models.Receipt{
		PurchaseTime: "16:05",
	}

	desc := rule.DescribePoints(&receipt, true)
	assert.Equal(t, Default_PurchaseTimeIsInRange_Points, desc.Points, "%v", desc.GetDescription())
}
