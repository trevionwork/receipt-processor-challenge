package rules

import (
	"testing"

	"github.com/fetch-rewards/receipt-processor-challenge/pkg/models"
	"github.com/stretchr/testify/assert"
)

func TestRetailerName(t *testing.T) {
	rule := NewDefaultRetailerName()
	receipt := models.Receipt{
		Retailer: "Mr. Test",
	}

	desc := rule.DescribePoints(&receipt, true)
	assert.Equal(t, desc.Points, Default_RetailerName_Each*6, "%v", desc.GetDescription())
}
