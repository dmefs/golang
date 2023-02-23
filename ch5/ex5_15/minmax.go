package ex515

func Min(x int, values ...int) int {
	for _, val := range values {
		if val < x {
			x = val
		}
	}
	return x
}

func Max(x int, values ...int) int {
	for _, val := range values {
		if val > x {
			x = val
		}
	}
	return x
}
