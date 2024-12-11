package rules

import (
	"fmt"
	"strings"

	"github.com/fetch-rewards/receipt-processor-challenge/pkg/models"
)

const TotalHaveNoCents_Default_Points = 50

type totalHaveNoCentsRule struct {
	Points int
}

func NewTotalHaveNoCents(points int) PointComputationRule {
	return &totalHaveNoCentsRule{Points: points}
}
func NewDefaultTotalHaveNoCents() PointComputationRule {
	return NewTotalHaveNoCents(TotalHaveNoCents_Default_Points)
}

func (rule *totalHaveNoCentsRule) Name() string {
	return "TotalHaveNoCents"
}
func (rule *totalHaveNoCentsRule) Description() string {
	return fmt.Sprintf("%d points if the total is a round dollar amount with no cents", rule.Points)
}
func (rule *totalHaveNoCentsRule) ComputePoints(receipt *models.Receipt) int {
	if strings.HasSuffix(receipt.Total, ".00") {
		return rule.Points
	}
	return 0
}
func (rule *totalHaveNoCentsRule) DescribePoints(receipt *models.Receipt, includeDescForZeroPoints bool) models.PointsRuleDescription {
	points := rule.ComputePoints(receipt)
	if points != 0 {
		return models.NewPointsRuleDescriptionf(
			points,
			"%d points because %s is a round dollar amount with no cents",
			rule.Points, receipt.Total,
		)
	} else if includeDescForZeroPoints {
		return models.NewPointsRuleDescriptionf(
			0,
			"no points because total price %s is not a round dollar and have some cents",
			receipt.Total,
		)
	} else {
		return models.PointsRuleDescription{}
	}
}
