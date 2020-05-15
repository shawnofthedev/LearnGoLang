package main

import (
	"fmt"
)

func main() {
	fmt.Println("Welcome Gophers :)")
	//fmt.Println("What would you like to do?")
	//mean()
	//exampleIf()
	//exampleSwitch()
	exampleFor()
}

func mean() {
	x, y := 1.0, 2.0

	fmt.Printf("x = %v, type of %T\n", x, x)
	fmt.Printf("y = %v, type of %T\n", y, y)

	mean := (x + y) / 2.0
	fmt.Printf("mean = %v, type of %T\n", mean, mean)
}

func exampleIf() {
	x := 10

	if x > 5 {
		fmt.Println("x is big")
	}

	if x > 100 {
		fmt.Println("x is very big")
	} else {
		fmt.Println("x is not that big")
	}

	if x > 5 && x < 15 {
		fmt.Println("x is just right")
	}

	if x < 20 || x > 30 {
		fmt.Println("x is out of range")
	}

	a := 11.0
	b := 20.0

	if frac := a / b; frac > 0.5 {
		fmt.Println("a is more than half of b")
	}
}

func exampleSwitch() {
	x := 2

	switch x {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	default:
		fmt.Printf("many")
	}

	//naked switch
	switch {
	case x > 100:
		fmt.Println("x is very big")
	case x > 10:
		fmt.Printf("x is big")
	default:
		fmt.Printf("x is small")
	}
}

func exampleFor() {
	for i := 0; i < 3; i++ {
		fmt.Println(i)
	}

	fmt.Println("----")
	for i := 0; i < 3; i++ {
		if i > 1 {
			break
		}
		fmt.Println(i)
	}

	fmt.Println("----")
	for i := 0; i < 3; i++ {
		if i < 1 {
			continue
		}
		fmt.Println(i)
	}
}
