package rules

import (
	"fmt"
	"unicode"

	"github.com/fetch-rewards/receipt-processor-challenge/pkg/models"
)

const Default_RetailerName_Each = 1

type retailerNameRule struct {
	IsValidCharacter   func(rune) bool
	ValidCharacterName string
	Each               int
}

func NewRetailerName(isValidChar func(rune) bool, validCharName string, each int) PointComputationRule {
	return &retailerNameRule{
		IsValidCharacter:   isValidChar,
		ValidCharacterName: validCharName,
		Each:               each,
	}
}
func NewDefaultRetailerName() PointComputationRule {
	return NewRetailerName(
		func(r rune) bool {
			return unicode.IsLetter(r) || unicode.IsNumber(r)
		},
		"alphanumeric",
		Default_RetailerName_Each,
	)
}

func (rule *retailerNameRule) Name() string {
	return "RetailerName"
}
func (rule *retailerNameRule) Description() string {
	return fmt.Sprintf("%d point for each %s character in retailer name", rule.Each, rule.ValidCharacterName)
}
func (rule *retailerNameRule) ComputePoints(receipt *models.Receipt) int {
	points := 0
	for _, r := range receipt.Retailer {
		if rule.IsValidCharacter(r) {
			points += rule.Each
		}
	}
	return points
}
func (rule *retailerNameRule) DescribePoints(receipt *models.Receipt, includeDescForZeroPoints bool) models.PointsRuleDescription {
	points := rule.ComputePoints(receipt)
	if points != 0 {
		return models.NewPointsRuleDescriptionf(
			points,
			"retailer name has %d %s charater(s) %d points each",
			points, rule.ValidCharacterName, rule.Each,
		)
	} else if includeDescForZeroPoints {
		return models.NewPointsRuleDescriptionf(
			0,
			"retailer name has no %s character",
			rule.ValidCharacterName,
		)
	} else {
		return models.PointsRuleDescription{}
	}
}
