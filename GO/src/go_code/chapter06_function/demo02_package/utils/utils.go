package utils

func Cal(n1 float64, n2 float64, operator byte) float64 {
	switch operator {
	case '+':
		return n1 + n2
	case '-':
		return n1 - n2
	case '*':
		return n1 * n2
	case '/':
		return n1 / n2
	default:
		return 0
	}
}
