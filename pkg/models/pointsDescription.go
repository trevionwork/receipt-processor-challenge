package models

import (
	"fmt"
	"strings"
)

type PointsRuleDescription struct {
	Points      int      `json:"points"`
	Description []string `json:"description"`
}

func NewPointsRuleDescription(points int, desc ...string) PointsRuleDescription {
	return PointsRuleDescription{Points: points, Description: desc}
}
func NewPointsRuleDescriptionf(points int, format string, args ...any) PointsRuleDescription {
	return NewPointsRuleDescription(points, fmt.Sprintf(format, args...))
}

func (desc *PointsRuleDescription) GetDescription() string {
	return strings.Join(desc.Description, "\n")
}

type PointsDescription struct {
	TotalPoints int                     `json:"totalPoints"`
	Details     []PointsRuleDescription `json:"details"`
}
