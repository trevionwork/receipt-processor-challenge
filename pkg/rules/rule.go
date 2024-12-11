package rules

import "github.com/fetch-rewards/receipt-processor-challenge/pkg/models"

type PointComputationRule interface {
	Name() string
	Description() string
	ComputePoints(receipt *models.Receipt) int
	DescribePoints(receipt *models.Receipt, includeDescForZeroPoints bool) models.PointsRuleDescription
}
