package rules

import (
	"fmt"

	"github.com/fetch-rewards/receipt-processor-challenge/pkg/models"
)

const Default_PointsForEveryNItems_N = 2
const Default_PointsForEveryNItems_Each = 5

type pointsForEveryNItemsRule struct {
	N    int
	Each int
}

func NewPointsForEveryNItems(n, each int) PointComputationRule {
	return &pointsForEveryNItemsRule{N: n, Each: each}
}
func NewDefaultPointsForEveryNItems() PointComputationRule {
	return NewPointsForEveryNItems(Default_PointsForEveryNItems_N, Default_PointsForEveryNItems_Each)
}

func (rule *pointsForEveryNItemsRule) Name() string {
	return "PointsForEveryNItems"
}
func (rule *pointsForEveryNItemsRule) Description() string {
	return fmt.Sprintf("%d points for every %d items on the receipt", rule.N, rule.Each)
}
func (rule *pointsForEveryNItemsRule) ComputePoints(receipt *models.Receipt) int {
	count := len(receipt.Items) / rule.N
	return count * rule.Each
}
func (rule *pointsForEveryNItemsRule) DescribePoints(receipt *models.Receipt, includeDescForZeroPoints bool) models.PointsRuleDescription {
	count := len(receipt.Items) / rule.N
	if count != 0 {
		return models.NewPointsRuleDescriptionf(
			count*rule.Each,
			"%d items (%d, %d items @ %d points each)",
			len(receipt.Items),
			count, rule.N,
			rule.Each,
		)
	} else if includeDescForZeroPoints {
		return models.NewPointsRuleDescription(
			0,
			"no points because there is not enough items in the list",
		)
	} else {
		return models.PointsRuleDescription{}
	}
}
