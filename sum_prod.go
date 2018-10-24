package main

import (
	"fmt"
	"strings"
)

func main() {
	var n, res uint64
	fmt.Println("Please enter a number")
	fmt.Scanf("%d", &n)

	var answer = "Y"
	fmt.Println("Do you want me to compute sum [Y/n] ? Otherwise product")
	fmt.Scanf("%s", &answer)

	sum := false
	if strings.ToLower(answer) == "y" {
		sum = true
	}

	if sum {
		for i := uint64(1); i <= n; i++ {
			res += i
		}
	} else {
		res = 1
		for i := uint64(1); i <= n; i++ {
			res *= i
		}
	}

	fmt.Println("sum =", res)
}
