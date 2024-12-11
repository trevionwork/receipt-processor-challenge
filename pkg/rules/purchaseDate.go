package rules

import (
	"fmt"
	"strconv"

	"github.com/fetch-rewards/receipt-processor-challenge/pkg/models"
)

const Default_PurchaseDate_Points = 6

type purchaseDateRule struct {
	IsOk   func(string) bool
	OkDesc string
	Points int
}

func NewPurchaseDate(isOk func(string) bool, okDesc string, points int) PointComputationRule {
	return &purchaseDateRule{
		IsOk:   isOk,
		OkDesc: okDesc,
		Points: points,
	}
}
func NewDefaultPurchaseDate() PointComputationRule {
	return NewPurchaseDate(
		func(s string) bool {
			day, _ := strconv.ParseInt(s[len(s)-2:], 10, 32)
			return day%2 == 1
		},
		"in an odd day",
		Default_PurchaseDate_Points,
	)
}

func (rule *purchaseDateRule) Name() string {
	return "PurchaseDate"
}
func (rule *purchaseDateRule) Description() string {
	return fmt.Sprintf("%d points if purchase date is %s", rule.Points, rule.OkDesc)
}
func (rule *purchaseDateRule) ComputePoints(receipt *models.Receipt) int {
	if rule.IsOk(receipt.PurchaseDate) {
		return rule.Points
	}
	return 0
}
func (rule *purchaseDateRule) DescribePoints(receipt *models.Receipt, includeDescForZeroPoints bool) models.PointsRuleDescription {
	points := rule.ComputePoints(receipt)
	if points != 0 {
		return models.NewPointsRuleDescriptionf(
			points,
			"%d points because purchase date is %s",
			rule.Points, rule.OkDesc,
		)
	} else if includeDescForZeroPoints {
		return models.NewPointsRuleDescriptionf(
			0,
			"no points because purchase date is not %s",
			rule.OkDesc,
		)
	} else {
		return models.PointsRuleDescription{}
	}
}
