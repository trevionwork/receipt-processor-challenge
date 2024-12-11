package rules

import (
	"fmt"
	"strings"

	"github.com/fetch-rewards/receipt-processor-challenge/pkg/models"
)

const Default_PurchaseTimeIsInRange_Start = "14:00"
const Default_PurchaseTimeIsInRange_End = "16:00"
const Default_PurchaseTimeIsInRange_Points = 10

type purchaseTimeIsInRangeRule struct {
	Start  string
	End    string
	Points int
}

func NewPurchaseTimeIsInRange(start, end string, points int) PointComputationRule {
	return &purchaseTimeIsInRangeRule{
		Start:  start,
		End:    end,
		Points: points,
	}
}
func NewDefaultPurchaseTimeIsInRange() PointComputationRule {
	return NewPurchaseTimeIsInRange(
		Default_PurchaseTimeIsInRange_Start,
		Default_PurchaseTimeIsInRange_End,
		Default_PurchaseTimeIsInRange_Points,
	)
}

func (rule *purchaseTimeIsInRangeRule) Name() string {
	return "PurchaseTimeIsInRange"
}
func (rule *purchaseTimeIsInRangeRule) Description() string {
	return fmt.Sprintf(
		"%d points if the time of purchase is after %s and before %s",
		rule.Points, rule.Start, rule.End)
}
func (rule *purchaseTimeIsInRangeRule) ComputePoints(receipt *models.Receipt) int {
	if n := strings.Compare(receipt.PurchaseTime, rule.Start); n < 0 {
		return 0
	}
	if n := strings.Compare(receipt.PurchaseTime, rule.End); n >= 0 {
		return 0
	}
	return rule.Points
}
func (rule *purchaseTimeIsInRangeRule) DescribePoints(receipt *models.Receipt, includeDescForZeroPoints bool) models.PointsRuleDescription {
	if n := strings.Compare(receipt.PurchaseTime, rule.Start); n < 0 {
		if includeDescForZeroPoints {
			return models.NewPointsRuleDescriptionf(
				0,
				"`%s` is before `%s`",
				receipt.PurchaseTime, rule.Start,
			)
		}
		return models.PointsRuleDescription{}
	}
	if n := strings.Compare(receipt.PurchaseTime, rule.End); n >= 0 {
		if includeDescForZeroPoints {
			return models.NewPointsRuleDescriptionf(
				0,
				"`%s` is after `%s`",
				receipt.PurchaseTime, rule.Start,
			)
		}
		return models.PointsRuleDescription{}
	}

	return models.NewPointsRuleDescriptionf(
		rule.Points,
		"`%s` is between `%s` and `%s`",
		receipt.PurchaseTime, rule.Start, rule.End,
	)
}
