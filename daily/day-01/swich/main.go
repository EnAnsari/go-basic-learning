package main

import "fmt"

// "fmt"

func comp(num1 int, num2 int) (isEqual bool, difference int) {
	switch {
	case num1 > num2:
		isEqual = false
		difference = num1 - num2
		// fallthrough
	case num2 > num1:
		isEqual = false
		difference = num2 - num1
	default:
		isEqual = true
		difference = 0
	}
	return
}

var x int8 = 45

func main() {
	// fmt.Println(comp(24, 52))

	// switch y := 5; x {
	// case 1:
	// 	fmt.Println("this number is one")
	// case 2:
	// 	fmt.Println("this number is two")
	// case 3:
	// 	fmt.Println("this number is three")
	// default:
	// 	fmt.Println("this number is not in range 1 to 3 and y =", y)
	// }

	// for i := 0; i < 10; i++ {
	// 	fmt.Println(i + 1)
	// }

	i := 0
	for i < 10 {
		fmt.Println(i + 1)
		i++
	}
}