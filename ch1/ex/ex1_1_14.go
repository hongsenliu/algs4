package ex

func Lg(N int) int {
	if N < 2 {
		return 0
	}
	var count int
	for n := N; n > 1; n = n / 2 {
		count++
	}
	return count
}
