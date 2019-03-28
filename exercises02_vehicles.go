package main

import (
	"fmt"
)

type vehicle struct {
	doors int
	color string
}

type truck struct {
	vehicle
	fourWheel bool
}

type sedan struct {
	vehicle
	luxury bool
}

type transportation interface {
	transportationDevice() string
}

func report(tran transportation) {
	fmt.Println(tran.transportationDevice())
}

func (t truck) transportationDevice() (string) {
	return fmt.Sprintln("this car goes brrrrrmmmmmmm")
}

func (s sedan) transportationDevice() (string) {
	return fmt.Sprintln("this car goes zoom zoom zoom")
}

func main() {
	t1 := truck{vehicle{
						2, 
						"red",
					},
				 	true,
				}

	s1 := sedan{vehicle{
						4, 
						"brown",
					}, 
					false,
				}

	fmt.Println(t1)
	fmt.Println(s1)

	report(s1)
	report(t1)
}