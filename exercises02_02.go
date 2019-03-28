package main

import "fmt"

type person struct {
	fName   string
	lName   string
	favFood []string
}

func (p person) walk() string {
	return fmt.Sprintln(p.fName, "is walking")
}

func main() {
	p1 := person{"grandma", "cereceres", []string{"tacos"}}

	p1.favFood = append(p1.favFood, "hamburgers", "cheese")

	fmt.Println(p1)
	fmt.Println(p1.lName)

	for _, v := range p1.favFood {
		fmt.Println(v)
	}

	s := p1.walk()

	fmt.Println(s)
}
