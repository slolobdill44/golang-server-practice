package main

import "fmt"

type gator int

func (g gator) greeting() () {
	fmt.Printf("Hello, I am a gator\n")
}

type flamingo bool

func (f flamingo) greeting() () {
	fmt.Printf("Hello, I am pink and beautiful and wonderful.\n")
}

type swampCreature interface {
	greeting()
}

func bayou(s swampCreature) () {
	s.greeting()
}

func main() {
	var g1 gator
	g1 = 42
	fmt.Println(g1)
	fmt.Printf("%T\n", g1)

	var x int

	fmt.Println(x)
	fmt.Printf("%T\n", x)

	x = int(g1)

	fmt.Println(x)

	g1.greeting()

	var f1 flamingo
	bayou(g1)
	bayou(f1)

}