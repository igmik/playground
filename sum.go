package main

import (
	"fmt"
	//"math/big"
)

func specSum(n uint64) uint64 {

	//sum := big.NewInt(0)

	n3 := uint64(n / 3)
	n5 := uint64(n / 5)
	n15 := uint64(n / 15)

	sum := 3 * n3 * (n3 + 1) / 2
	sum += 5 * n5 * (n5 + 1) / 2
	sum -= 15 * n15 * (n15 + 1) / 2

	return sum
}

func main() {
	var n, sum uint64
	fmt.Println("Please enter a number")
	fmt.Scanf("%d", &n)

	fmt.Println("n =", n)

	for i := uint64(1); i <= n; i++ {
		sum += i
	}
	fmt.Println("sum =", sum)

	fmt.Println(specSum(uint64(17)))
}
