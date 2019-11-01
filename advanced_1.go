package main

import (
	"fmt"
)

func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func PrintTab(s1 string, s2 string, tab [][]int) {
    fmt.Print(" ")
    for j:=0; j<len(s2); j++ {
        fmt.Printf(" %c",s2[j])
    }
    fmt.Println()
    for i:=0; i<len(s1); i++ {
        fmt.Printf("%c", s1[i])
        fmt.Println(tab[i+1][1:])
    }
}

func main() {
    var s1, s2, ss1, ss2 string
    var l1, l2 int
	fmt.Println("Please enter a first string")
	fmt.Scanf("%s", &ss1)
    //l1 = len(s1)
	fmt.Println("Please enter a second string")
	fmt.Scanf("%s", &ss2)
    //l2 = len(s2)

    if len(ss2) > len(ss1) {
        s1 = ss2
        s2 = ss1
        l1 = len(ss2)
        l2 = len(ss1)
    } else {
        s1 = ss1
        s2 = ss2
        l1 = len(ss1)
        l2 = len(ss2)
    }

    tab := make([][]int, l1+1)
    for i := range tab {
        tab[i] = make([]int, l2+1)
    }

    // Mark duplicates in a table
    var res []string
    var ind = 1
    for j:=0; j<=l2; j++ {
        for i:=0; i<=l1; i++ {
            if i == 0 || j == 0 {
                tab[i][j] = 0
            } else if s1[i-1] == s2[j-1] {
                tab[i][j] = tab[i-1][j-1] + 1
                if tab[i][j] == ind {
                    res = append(res, string(s1[i-1]))
                    ind += 1
                }
            } else {
                 tab[i][j] = Max(tab[i-1][j] , tab[i][j-1]) 
            }
        }
    }
    //PrintTab(s1, s2, tab)

    fmt.Println("Longest common subsequence is", res)
}
