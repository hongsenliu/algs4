package ex

func Histogram(a []int, m int) []int {
	if a == nil || len(a) == 0 || m <= 0 {
		return nil
	}

	h := make([]int, m)
	for _, i := range a {
		h[i]++
	}
	return h
}
