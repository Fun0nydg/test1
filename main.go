package test

import "fmt"

var (
	a int = 1
)

func max(a, b int) int {
	if a < b {
		return b
	} else {
		return a
	}
	fmt.Print("success")
}
