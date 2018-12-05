package main

import "fmt"

//import "math/rand"
//import "time"
//import "container/list"
//import "reflect"

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func PrintContent(l []int) {
	for _, a := range l {
		fmt.Println(a)
	}
}

func Number2List(num int) []int {
	mylist := []int{}

	var digit int

	for num > 0 {
		digit = num / 10
		digit = num - digit*10
		num = num / 10
		mylist = append(mylist, digit)
	}

	return mylist
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

	// 11 Two sorted lists
	fmt.Println("\nConvert number to list:")
	num := 123456
	PrintContent(Number2List(num))

}
