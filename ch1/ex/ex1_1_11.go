package ex

import "fmt"

func twoDBool(b [][]bool) {
	for _, r := range b {
		for _, e := range r {
			if e {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Print("\n")
	}
}
