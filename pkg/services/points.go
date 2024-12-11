package services

import (
	"github.com/fetch-rewards/receipt-processor-challenge/pkg/models"
	"github.com/fetch-rewards/receipt-processor-challenge/pkg/rules"
)

type PointsService struct {
	rules []rules.PointComputationRule
}

func NewPointsService(rules ...rules.PointComputationRule) *PointsService {
	return &PointsService{rules: rules}
}
func NewDefaultPointsService() *PointsService {
	return NewPointsService(
		rules.NewDefaultRetailerName(),
		rules.NewDefaultTotalHaveNoCents(),
		rules.NewDefaultTotalIsAMultipleOfNCents(),
		rules.NewDefaultPointsForEveryNItems(),
		rules.NewDefaultItemDesciptionLength(),
		rules.NewDefaultPurchaseDate(),
		rules.NewDefaultPurchaseTimeIsInRange(),
	)
}

func (svc *PointsService) ComputePoints(receipt *models.Receipt) int {
	points := 0
	for _, rule := range svc.rules {
		rulePoints := rule.ComputePoints(receipt)
		points += rulePoints
	}

	return points
}
func (svc *PointsService) DescribePoints(receipt *models.Receipt, includeDescForZeroPoints bool) models.PointsDescription {
	var desc models.PointsDescription
	for _, rule := range svc.rules {
		ruleDesc := rule.DescribePoints(receipt, includeDescForZeroPoints)
		if ruleDesc.Points != 0 || includeDescForZeroPoints {
			desc.Details = append(desc.Details, ruleDesc)
		}
		desc.TotalPoints += ruleDesc.Points
	}

	return desc
}
