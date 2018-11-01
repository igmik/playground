package main

import "fmt"

import "math/rand"
import "time"
import "container/list"
import "reflect"

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

func PrintContent(l *list.List) {
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}

func LargestElement(l *list.List) int {
	max := int(-1)
	for e := l.Front(); e != nil; e = e.Next() {
		max = Max(e.Value.(int), max)
	}
	return max
}

func ReverseList(l *list.List) *list.List {
	front := l.Front()
	for i, e := 0, l.Back(); i < l.Len(); i, e = i+1, l.Back() {
		l.MoveBefore(e, front)
	}
	return l
}

func IsInList(v int, l *list.List) bool {
	for e := l.Front(); e != nil; e = e.Next() {
		if e.Value == v {
			return true
		}
	}
	return false
}

func OddElements(l *list.List) *list.List {
	newList := list.New()
	for i, e := 1, l.Front(); i <= l.Len(); i, e = i+1, e.Next() {
		if i%2 != 0 {
			newList.PushBack(e.Value)
		}
	}
	return newList
}

func Adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func IsPolindrome(a string) bool {
	l := len(a) - 1
	for i := 0; i < len(a)/2; i++ {
		if a[i] != a[l-i] {
			return false
		}
	}
	return true
}

func ConcatAlt(l1, l2 *list.List) *list.List {
	for e1, e2 := l1.Front(), l2.Front(); (e1 != nil) && (e2 != nil); e1, e2 = e1.Next(), e2.Next() {
		l1.InsertAfter(e2.Value, e1)
		e1 = e1.Next()
	}
	return l1
}

func main() {
	src := rand.NewSource(time.Now().UnixNano())
	rnd := rand.New(src)

	myList := list.New()
	fmt.Println(reflect.TypeOf(myList.Front()))
	//myLen := rnd.Intn(100)
	myLen := 5

	for i := 0; i < myLen; i++ {
		myList.PushBack(rnd.Intn(10000))
	}

	fmt.Println("Content of list:")
	PrintContent(myList)

	// 1 Largest element
	fmt.Println("\nMaximum value in list:")
	fmt.Println(LargestElement(myList))

	// 2 Reversed list
	fmt.Println("\nReversed list:")
	PrintContent(ReverseList(myList))

	// 3 is value in list
	n := rnd.Int()
	fmt.Println("\nCheck if value", n, "in the list:")
	fmt.Println(IsInList(n, myList))

	// 4 elements on odd position
	fmt.Println("\nElements on odd position:")
	PrintContent(OddElements(myList))

	// 5 running total
	fmt.Println("\nResults of running total:")
	sum := Adder()
	for e := myList.Front(); e != nil; e = e.Next() {
		fmt.Println(sum(e.Value.(int)))
	}

	// 6 is string a polindromi
	fmt.Println("\nFind if a word is a polindrome:")
	arr := [6]string{"kayak", "boat", "anna", "level", "mom", "golang"}
	myList2 := list.New()
	for _, a := range arr {
		fmt.Println("is", a, "a polindrome:", IsPolindrome(a))
		myList2.PushBack(a)
	}

	// 8 Print perfect squares
	fmt.Println("\nPrint perfect squares:")
	//OnAll()

	// 9 Concatenate 2 lists
	fmt.Println("\nConcatenated list:")
	myList.PushBackList(myList2)
	PrintContent(myList)

	// 10 Concatenation alternatingly
	fmt.Println("\nConcatenated alternatingly list:")
	PrintContent(ConcatAlt(myList, myList2))

}
