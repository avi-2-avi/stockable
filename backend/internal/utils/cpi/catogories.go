package utils

func CategorizeCPI(normalizedCPI float64) string {
	switch {
	case normalizedCPI >= 90:
		return "+ Strong Buy"
	case normalizedCPI >= 80:
		return "Strong Buy"
	case normalizedCPI >= 70:
		return "Buy"
	case normalizedCPI >= 60:
		return "+ Hold"
	case normalizedCPI >= 50:
		return "Hold"
	case normalizedCPI >= 40:
		return "- Hold"
	case normalizedCPI >= 30:
		return "Sell"
	case normalizedCPI >= 20:
		return "Strong Sell"
	default:
		return "Avoid"
	}
}
