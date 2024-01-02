package internal

import "math"

func Minimum(nums ...int) int {
	min := math.Inf(1)
	for _, num := range nums {
		min = math.Min(min, float64(num))
	}
	return int(min)
}

func Average(nums ...int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	return sum / len(nums)
}

func Maximum(nums ...int) int {
	max := math.Inf(-1)
	for _, num := range nums {
		max = math.Max(max, float64(num))
	}
	return int(max)
}
