package main

import (
	"fmt"
	"strings"
)

func main() {

	N := 1000
	fmt.Println("Think a number from 1 to ", N)

	ming := 0
	maxg := N
	guess := N
	greater := false
	for {

		if greater {
			//guess += int(float32((maxg-guess)/2) + 0.5)
			guess += (maxg - guess) / 2
		} else {
			//guess -= int(float32((guess-ming)/2) + 0.5)
			guess -= (guess - ming) / 2
		}

		answer := "Y"
		fmt.Println("Is the number greater than", guess, "[Y/n]")
		fmt.Scanf("%s", &answer)

		greater = false
		if strings.ToLower(answer) == "y" {
			greater = true
			ming = guess
		} else {
			greater = false
			maxg = guess
		}

		if maxg-ming == 1 {
			break
		}

		fmt.Println(ming, maxg)
	}

	fmt.Println("Your number is ", maxg)

}
