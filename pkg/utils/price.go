package utils

import "strconv"

func PriceToCents(s string) int {
	price, _ := strconv.ParseFloat(s, 64)
	return int(price * 100)
}
