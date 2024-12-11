package rules

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/fetch-rewards/receipt-processor-challenge/pkg/models"
)

const Default_ItemDesciptionLength_N = 3
const Default_ItemDesciptionLength_PriceMultuplier = 0.2

type itemDesciptionLengthRule struct {
	N               int
	PriceMultiplier float64
}

func NewItemDesciptionLength(n int, priceMultiplier float64) PointComputationRule {
	return &itemDesciptionLengthRule{N: n, PriceMultiplier: priceMultiplier}
}
func NewDefaultItemDesciptionLength() PointComputationRule {
	return NewItemDesciptionLength(Default_ItemDesciptionLength_N, Default_ItemDesciptionLength_PriceMultuplier)
}

func (rule *itemDesciptionLengthRule) Name() string {
	return "ItemDesciptionLength"
}
func (rule *itemDesciptionLengthRule) Description() string {
	return fmt.Sprintf(
		"If the trimmed length of the item description is a multiple of %d, multiply the price by `%v` and round up to the nearest integer. The result is the number of points earned",
		rule.N, rule.PriceMultiplier)
}
func (rule *itemDesciptionLengthRule) ComputePoints(receipt *models.Receipt) int {
	points := 0
	for i := range receipt.Items {
		item := &receipt.Items[i]
		itemDesc := strings.TrimSpace(item.ShortDescription)
		if (len(itemDesc) % rule.N) == 0 {
			itemPrice, _ := strconv.ParseFloat(item.Price, 64)
			itemPoints := itemPrice * rule.PriceMultiplier
			roundupItemPoint := int(math.Ceil(itemPoints))
			points += roundupItemPoint
		}
	}
	return points
}
func (rule *itemDesciptionLengthRule) DescribePoints(receipt *models.Receipt, includeDescForZeroPoints bool) models.PointsRuleDescription {
	desc := models.PointsRuleDescription{}
	for i := range receipt.Items {
		item := &receipt.Items[i]
		itemDesc := strings.TrimSpace(item.ShortDescription)
		if (len(itemDesc) % rule.N) == 0 {
			itemPrice, _ := strconv.ParseFloat(item.Price, 64)
			itemPoints := itemPrice * rule.PriceMultiplier
			roundupItemPoint := int(math.Ceil(itemPoints))
			desc.Description = append(
				desc.Description,
				fmt.Sprintf(
					"%q has %d character (a multiple of %d), item price %s * %v = %.2f, round up to %d",
					itemDesc, len(itemDesc), rule.N,
					item.Price, rule.PriceMultiplier, itemPoints,
					roundupItemPoint,
				),
			)
			desc.Points += roundupItemPoint
		} else if includeDescForZeroPoints {
			desc.Description = append(
				desc.Description,
				fmt.Sprintf(
					"%q is %d character which is not a multiple of %d",
					itemDesc, len(itemDesc), rule.N,
				),
			)
		}
	}
	return desc
}
