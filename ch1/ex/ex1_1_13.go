package ex

func TranM(ar [][]int) [][]int {
	if ar == nil || len(ar) == 0 {
		return nil
	}
	m := len(ar)
	n := len(ar[0])
	tm := make([][]int, n)
	for i := range tm {
		tm[i] = make([]int, m)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			tm[j][i] = ar[i][j]
		}
	}
	return tm
}
