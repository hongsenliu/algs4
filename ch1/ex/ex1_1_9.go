package ex

import "strconv"

func intToBitStr(n int) string {
	var s string
	for i := n; i > 0; i = i / 2 {
		a := strconv.Itoa(i % 2)
		s = a + s
	}
	return s
}
