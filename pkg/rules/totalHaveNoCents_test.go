package rules

import (
	"testing"

	"github.com/fetch-rewards/receipt-processor-challenge/pkg/models"
	"github.com/stretchr/testify/assert"
)

func TestTotalHaveNoCents(t *testing.T) {
	rule := NewDefaultTotalHaveNoCents()
	receipt := models.Receipt{
		Total: "12.00",
	}

	desc := rule.DescribePoints(&receipt, true)
	assert.Equal(t, TotalHaveNoCents_Default_Points, desc.Points, "%v", desc.GetDescription())
}
func TestTotalHaveNoCentsNotRound(t *testing.T) {
	rule := NewDefaultTotalHaveNoCents()
	receipt := models.Receipt{
		Total: "12.01",
	}

	desc := rule.DescribePoints(&receipt, true)
	assert.Equal(t, 0, desc.Points, "%v", desc.GetDescription())
}
