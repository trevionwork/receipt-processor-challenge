package rules

import (
	"math"
	"testing"

	"github.com/fetch-rewards/receipt-processor-challenge/pkg/models"
	"github.com/stretchr/testify/assert"
)

func TestItemDesciptionLengthRule(t *testing.T) {
	rule := NewDefaultItemDesciptionLength()
	receipt := models.Receipt{
		Retailer: "Target",
		Items: []models.ReceiptItem{
			{ShortDescription: "  1   ", Price: "1.49"},
			{ShortDescription: "  12  ", Price: "2.12"},
			{ShortDescription: "  123 ", Price: "3.25"},
			{ShortDescription: "  1234   ", Price: "4.26"},
			{ShortDescription: "  12345  ", Price: "5.26"},
			{ShortDescription: "  123456 ", Price: "6.26"},
		},
	}
	desc := rule.DescribePoints(&receipt, true)
	// `123` is a multiple of 3
	expected := int(math.Ceil(3.25 * Default_ItemDesciptionLength_PriceMultuplier))
	// `123456` is a multiple of 3
	expected += int(math.Ceil(6.26 * Default_ItemDesciptionLength_PriceMultuplier))
	assert.Equal(t, expected, desc.Points, "%v", desc.GetDescription())
}
