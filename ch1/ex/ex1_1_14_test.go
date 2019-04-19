package ex

import "testing"

func TestLg(t *testing.T) {
	ts := []struct {
		input int
		want  int
	}{
		{1, 0},
		{2, 1},
		{16, 4},
		{25, 4},
	}
	for _, c := range ts {
		if a := Lg(c.input); a != c.want {
			t.Errorf("want %d, actual %d", c.want, a)
		}
	}
}
