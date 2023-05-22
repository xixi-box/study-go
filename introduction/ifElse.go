package introduction

import "fmt"

func IfElse() {
	if 7%2 == 0 {
		fmt.Println("偶数")
	} else {
		fmt.Println("奇数")
	}
	if 8%2 == 0 {
		fmt.Println("偶数")
	} else {
		fmt.Println("奇数")
	}
	num := 10
	if num < 0 {
		fmt.Println(num, "负数")
	} else if num < 10 {
		fmt.Println(num, "个位数")
	} else {
		fmt.Println(num, "多位数")
	}
}
