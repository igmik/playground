package main

import "fmt"

import "math"

//import "time"
//import "container/list"
//import "reflect"

func Min(x, y uint) uint {
	if x < y {
		return x
	}
	return y
}

func Max(x, y uint) uint {
	if x > y {
		return x
	}
	return y
}

func getDigits_base10(x int) uint {
	var ndigits uint

	if x == 0 {
		return 1
	}

	for x > 0 {
		ndigits++
		x = x / 10
	}

	return ndigits
}

func split_at(num int, digits uint) (int, int) {
	divisor := int(math.Pow(10, float64(digits)))

	if num >= divisor {
		return num / divisor, num % divisor
	} else {
		return 0, num
	}
}

func Karatsuba(x, y int) int {

	if x < 10 || y < 10 {
		return x * y
	}

	m := Min(getDigits_base10(x), getDigits_base10(y))
	m2 := m / 2

	x_high, x_low := split_at(x, m2)
	y_high, y_low := split_at(y, m2)

	z0 := Karatsuba(x_low, y_low)
	z1 := Karatsuba((x_low + x_high), (y_low + y_high))
	z2 := Karatsuba(x_high, y_high)

	return (z2 * int(math.Pow(10, float64(2*m2)))) + (z1-z2-z0)*int(math.Pow(10, float64(m2))) + z0

}

func main() {
	//src := rand.NewSource(time.Now().UnixNano())
	//rnd := rand.New(src)

	//myList := list.New()
	//fmt.Println(reflect.TypeOf(myList.Front()))
	////myLen := rnd.Intn(100)
	//myLen := 5

	//for i := 0; i < myLen; i++ {
	//	myList.PushBack(rnd.Intn(10000))
	//}

	//fmt.Println("Content of list:")
	//PrintContent(myList)

	// 15 Dot product using Karatsuba algorithm
	fmt.Println("\nDot product using Karatsuba algorithm")
	num1 := 123456
	num2 := 654321

	fmt.Println(Karatsuba(num1, num2))

}
