package ex

func tranM(ar [][]int) [][]int {
	if ar == nil || len(ar) == 0 {
		return nil
	}
	m := len(ar)
	n := len(ar[0])
	tm := [][]int{}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			tm[j][i] = ar[i][j]
		}
	}
	return tm
}
