package fib

import (
	"fmt"
	"testing"
)

func TestFibList(t *testing.T) {
	var a int = 1
	var b int = 1
	//var (
	// a int = 1
	// b int = 1
	//)
	for i := 0; i < 5; i++ {
		fmt.Printf(" ", b)
		tmp := a
		a = b
		b = tmp + a
	}
	fmt.Println()
}
