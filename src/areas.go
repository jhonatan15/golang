package main

import (
	"fmt"
)

func main() {
	// Rectangle
	const rSide = 20
	const rBase = 10
	fmt.Println("The area of rectangle is:", rSide*rBase)

	// Trapezium
	const tBase1 = 4
	const tBase2 = 10
	const tHeight = 4
	fmt.Println("The area of trapezium is:", ((tBase1+tBase2)*tHeight)/2)

	// Circle
	const cDiameter1 = 38 / 2
	const pi3 = 3.14
	fmt.Println("The area of trapezium is:", (cDiameter1*cDiameter1)*pi3)
}
