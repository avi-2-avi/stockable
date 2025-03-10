package utils

import "math"

func CalculateActionImpactScore(action string) float64 {
	switch action {
	case "reiterated by":
		return 0.5
	case "target lowered by":
		return -1.0
	case "target raised by":
		return 1.0
	case "initiated by":
		return 1.5
	case "upgraded by":
		return 2.0
	case "downgraded by":
		return -2.0
	default:
		return 0
	}
}

func CalculateRatingChangeImpact(from string, to string) float64 {
	ratingImpact := map[string]float64{
		"Outperform":   4,
		"Strong-Buy":   4,
		"Buy":          3,
		"Overweight":   3,
		"Neutral":      1,
		"Hold":         1,
		"Equal Weight": 1,
		"Sell":         -1,
		"Underperform": -1,
		"Underweight":  -1,
	}

	if from == to {
		return ratingImpact[from]
	}

	if from == "Neutral" || from == "Hold" || from == "Equal Weight" {
		if to == "Strong-Buy" || to == "Buy" || to == "Outperform" || to == "Overweight" {
			return 2
		}
	}
	if from == "Sell" || from == "Underperform" || from == "Underweight" {
		if to == "Strong-Buy" || to == "Buy" || to == "Outperform" || to == "Overweight" {
			return 3
		}
	}
	if from == "Strong-Buy" || from == "Buy" || from == "Outperform" || from == "Overweight" {
		if to == "Neutral" || to == "Hold" || to == "Equal Weight" {
			return -2
		}
	}
	if from == "Strong-Buy" || from == "Buy" || from == "Outperform" || from == "Overweight" {
		if to == "Sell" || to == "Underperform" || to == "Underweight" {
			return -3
		}
	}
	if from == "Sell" || from == "Underperform" || from == "Underweight" {
		if to == "Strong-Buy" {
			return 4
		}
	}
	if from == "Strong-Buy" {
		if to == "Sell" || to == "Underperform" || to == "Underweight" {
			return -4
		}
	}

	return 0
}

func CalculateTargetAdjustment(from float64, to float64) float64 {
	if from == 0 {
		return 0
	}
	tap := (to - from) * 100.0 / from
	if tap > 100 {
		return 100
	} else if tap < -100 {
		return -100
	}
	return tap
}

func CalculateCPI(ais float64, rci float64, tap float64) float64 {
	rawCPI := (0.3 * ais) + (0.35 * rci) + (0.35 * tap)
	normalizedCPI := ((rawCPI - (-54.5)) / (37 - (-54.5))) * 100

	return math.Round(normalizedCPI*100) / 100
}
