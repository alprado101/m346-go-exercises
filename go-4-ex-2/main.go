package main

import (
	"fmt"
	"math"
)

// TODO: implement the function computeHypotenuse using math.Pow and math.Sqrt

func computeHypotenuse(a float64, b float64) float64 {
	// Formel: c = sqrt(a^2 + b^2)
	return math.Sqrt(math.Pow(a, 2) + math.Pow(b, 2))
}

func main() {
	// TODO: call the function computeHypotenuse

	h1 := computeHypotenuse(3, 4)
	fmt.Println("Hypotenuse von 3 und 4 ist:", h1) // 5

	h2 := computeHypotenuse(5, 12)
	fmt.Println("Hypotenuse von 5 und 12 ist:", h2) // 13

	fmt.Println(computeHypotenuse(8.0, 15.0)) // 17
}
