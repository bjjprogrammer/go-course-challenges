package main

import (
	"fmt"
	"math"
)

const PRESICION = 2

func main() {
	var side1, side2 float64
	fmt.Print("Enter the first side of the rectangle: ")
	fmt.Scanln(&side1)
	fmt.Print("Enter the second side of the rectangle: ")
	fmt.Scanln(&side2)

	area := side1 * side2 / 2
	hipotenusa := math.Sqrt(math.Pow(side1, 2) + math.Pow(side2, 2))
	perimeter := side1 + side2 + hipotenusa
	fmt.Println("The hypotenuse of the rectangle is: ", hipotenusa)
	fmt.Printf("The area of the rectangle is: %.*f\n", PRESICION, area)
	fmt.Printf("The perimeter of the rectangle is: %.*f\n", PRESICION, perimeter)
}
