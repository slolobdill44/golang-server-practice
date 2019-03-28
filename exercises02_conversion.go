package main

import "fmt"

func main() {

	var s string

	s = "i'm sorry dave i can't do that"

	fmt.Println(s)

	fmt.Println([]byte(s))
	
	fmt.Println([]byte(s))

	fmt.Println((s)[:14])

	for _, v := range []byte(s) {
		fmt.Println(string(v))
	}
}