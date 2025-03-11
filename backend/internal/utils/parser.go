package utils

import (
	"regexp"
	"strconv"
)

func ParsePrice(unparsedPrice string) float64 {
	priceRegex := regexp.MustCompile(`[-+]?[0-9]*\.?[0-9]+`)
	matchPrice := priceRegex.FindString(unparsedPrice)
	if matchPrice == "" {
		return 0
	}
	price, _ := strconv.ParseFloat(matchPrice, 64)
	return price
}
