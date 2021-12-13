package main

import (
	"fmt"
	"strconv"
)

func main() {
	for i := 1; i <= 100; i++ {
		ThreeFive(i)
	}
}

func ThreeFive(number int) {
	if number%3 == 0 && number%5 == 0 {
		fmt.Println("ThreeFive")
	} else if number%3 == 0 {
		fmt.Println("Three")
	} else if number%5 == 0 {
		fmt.Println("Five")
	} else {
		fmt.Println(strconv.Itoa(number))
	}
}
