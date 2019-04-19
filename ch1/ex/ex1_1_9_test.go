package ex

import (
	"fmt"
	"testing"
)

func TestIntToBitStr(t *testing.T) {
	fmt.Println(intToBitStr(2))
}

func TestTranM(t *testing.T) {
	m := [][]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Println(TranM(m))
}
