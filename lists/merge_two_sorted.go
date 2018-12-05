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

func Merge2Arrays(l1 []int, l2 []int) []int {
	merged := []int{}

	i, j := 0, 0

	for i < len(l1) && j < len(l2) {

		if l1[i] < l2[j] {
			merged = append(merged, l1[i])
			i = i + 1
		} else {
			merged = append(merged, l2[j])
			j = j + 1
		}
	}

	for i < len(l1) {
		merged = append(merged, l1[i])
		i = i + 1
	}

	for j < len(l2) {
		merged = append(merged, l2[j])
		j = j + 1
	}

	return merged
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
	fmt.Println("\nMerge 2 sorted lists:")
	sorted1 := []int{10, 20, 50, 150, 200, 250, 300}
	sorted2 := []int{30, 40, 45, 100, 400, 450}
	PrintContent(Merge2Arrays(sorted1, sorted2))

}
