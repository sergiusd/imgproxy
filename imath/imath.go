package imath

import "math"

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func MinNonZero(a, b int) int {
	switch {
	case a == 0:
		return b
	case b == 0:
		return a
	}

	return Min(a, b)
}

func Round(a float64) int {
	return int(math.Round(a))
}

func Scale(a int, scale float64) int {
	if a == 0 {
		return 0
	}

	return Round(float64(a) * scale)
}
