package rules

import (
	"testing"

	"github.com/fetch-rewards/receipt-processor-challenge/pkg/models"
	"github.com/stretchr/testify/assert"
)

func TestXxx(t *testing.T) {
	rule := NewDefaultTotalIsAMultipleOfNCents()
	tests := []struct {
		Receipt  models.Receipt
		Expected int
	}{
		{
			Receipt:  models.Receipt{Total: "1.25"},
			Expected: TotalIsAMultipleOfNCents_Default_Points,
		},
		{
			Receipt:  models.Receipt{Total: "12.50"},
			Expected: TotalIsAMultipleOfNCents_Default_Points,
		},
		{
			Receipt:  models.Receipt{Total: "12.75"},
			Expected: TotalIsAMultipleOfNCents_Default_Points,
		},
		{
			Receipt:  models.Receipt{Total: "125.00"},
			Expected: TotalIsAMultipleOfNCents_Default_Points,
		},
		{
			Receipt:  models.Receipt{Total: "125.20"},
			Expected: 0,
		},
	}

	for _, test := range tests {
		desc := rule.DescribePoints(&test.Receipt, true)
		assert.Equal(t, test.Expected, desc.Points, "%v", desc.GetDescription())
	}

}
