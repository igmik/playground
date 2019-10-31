package main

import (
	"fmt"
)

func main() {
    var s1, s2 string
    var l1, l2 int
	fmt.Println("Please enter a first string")
	fmt.Scanf("%s", &s1)
    l1 = len(s1)
	fmt.Println("Please enter a second string")
	fmt.Scanf("%s", &s2)
    l2 = len(s2)

    if l1 > l2 {
        tmp := l2
        l2 = l1
        l1 = tmp

        tmps := s2
        s2 = s1
        s1 = tmps
    }

    for l:=l1; l>0; l-- {
        for j:=0; j+l<=l1; j++ {  
            sub1 := s1[j:j+l]
            for i:=0; i+l<=l2; i++ {
                if s2[i:i+l] == sub1 {
                    fmt.Println("Longest common substring is", sub1)
                    return
                }
            }
        }
	}

    fmt.Println("No common substring found :-(")
}
