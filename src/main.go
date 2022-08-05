package main

import (
	"fmt"
)

func main() {
	// Declaration of constants
	const pi float64 = 3.14
	const pi2 = 3.14

	fmt.Println("pi:", pi)
	fmt.Println("pi2:", pi2)

	// Declaration of integer variables
	base := 12
	var height int = 14
	var area int

	fmt.Println(base, height, area)

	// Zero values
	var a int
	var b float64
	var c string
	var d bool

	fmt.Println(a, b, c, d)

	//Square area
	const squareBase = 10
	squareArea := squareBase * squareBase
	fmt.Println("Square area is:", squareArea)

	x := 10
	y := 50

	//Sum
	result := x + y
	fmt.Println("Sum:", result)

	//Substract
	result = y - x
	fmt.Println("Susbtract:", result)

	//Multiplication
	result = y * x
	fmt.Println("Multiplication:", result)

	//Division
	result = y / x
	fmt.Println("Division:", result)

	// Module
	result = y % x
	fmt.Println("Module:", result)

	// Increase
	x++
	fmt.Println("Increase:", x)

	// Decrease
	x--
	fmt.Println("Decrease:", x)

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
