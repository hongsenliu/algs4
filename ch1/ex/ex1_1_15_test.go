package ex

import (
	"fmt"
	"testing"
)

func TestHistoram(t *testing.T) {
	a := []int{0, 1, 2, 3, 2, 1, 2, 3}
	fmt.Println(Histogram(a, 4))
}
