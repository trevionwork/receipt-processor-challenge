package rules

import (
	"fmt"

	"github.com/fetch-rewards/receipt-processor-challenge/pkg/models"
	"github.com/fetch-rewards/receipt-processor-challenge/pkg/utils"
)

const TotalIsAMultipleOfNCents_Default_Cents = 25
const TotalIsAMultipleOfNCents_Default_Points = 25

type totalIsAMultipleOfNCentsRule struct {
	Cents  int
	Points int
}

func NewTotalIsAMultipleOfNCents(cents, points int) PointComputationRule {
	return &totalIsAMultipleOfNCentsRule{Cents: cents, Points: points}
}
func NewDefaultTotalIsAMultipleOfNCents() PointComputationRule {
	return NewTotalIsAMultipleOfNCents(TotalIsAMultipleOfNCents_Default_Cents, TotalIsAMultipleOfNCents_Default_Points)
}

func (rule *totalIsAMultipleOfNCentsRule) Name() string {
	return "totalIsAMultipleOfNCents"
}
func (rule *totalIsAMultipleOfNCentsRule) Description() string {
	return fmt.Sprintf("%d points if the total price is a multiple of %d cents", rule.Points, rule.Cents)
}
func (rule *totalIsAMultipleOfNCentsRule) ComputePoints(receipt *models.Receipt) int {
	cents := utils.PriceToCents(receipt.Total)
	if cents%rule.Cents == 0 {
		return rule.Points
	}
	return 0
}
func (rule *totalIsAMultipleOfNCentsRule) DescribePoints(receipt *models.Receipt, includeDescForZeroPoints bool) models.PointsRuleDescription {
	points := rule.ComputePoints(receipt)
	if points != 0 {
		return models.NewPointsRuleDescriptionf(
			points,
			"%d points because total price %s is a multiple of %d cents",
			rule.Points,
			receipt.Total,
			rule.Cents,
		)
	} else if includeDescForZeroPoints {
		return models.NewPointsRuleDescriptionf(
			0,
			"no points because total price %s if not a multiple of %d cents",
			receipt.Total, rule.Cents,
		)
	} else {
		return models.PointsRuleDescription{}
	}
}
