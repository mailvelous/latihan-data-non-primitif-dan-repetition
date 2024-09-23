package main

import (
	"fmt"
)

func main() {

	case3A()
	fmt.Println()
	fmt.Println()
	case3B()
	fmt.Println()
	fmt.Println()
	case3C()
	fmt.Println()
	fmt.Println()
	case3D()
	fmt.Println()
	fmt.Println()
	case3E()

}

func case3A() {
	var star string
	star = "*"

	for i := 0; i < 5; i++ {
		fmt.Printf("%s", star)
	}
}

func case3B() {
	for i := 1; i <= 10; i++ {
		if i%2 == 1 {
			fmt.Printf("%d" ,i)
		}
	}
}


func case3C() {
	for i := 1; i <= 5; i++ {
		fmt.Println()
		for i := 1; i <= 5; i++ {
			fmt.Printf("%d", i)
		}
	}

}

func case3D() {
	number := 1

	for i := 1; i <= 5; i++ {
		for i := 1; i <= 5; i++ {
			fmt.Printf("%d", number)
		}
		fmt.Println()
		number = number + 1
	}
}

func case3E() {
	var number int
	var space string

	number = 10
	space = " "

	for i := 1; i <= 5; i++ {
		for i := 1; i <= 5; i++ {
			number = number + 1
			fmt.Printf("%s", space)
			fmt.Printf("%d", number)
		}
		fmt.Println()
		number = number + 5
	}


}