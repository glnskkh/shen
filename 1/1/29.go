package shen

func CountSolutionsSlow(n int) int {
	count := 0

	for k := 0; k*k < n; k++ {
		for l := 0; l*l+k*k < n; l++ {
			count++
		}
	}

	return count
}
