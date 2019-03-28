package main

import (
	"fmt"
)



func main() {
	myslice := []int{3, 4, 34, 3}

	fmt.Println(myslice)

	for i, _ := range myslice {
		fmt.Println(i)
	}

	for i, v := range myslice {
		fmt.Println(i, v)
	}
}