package main

import (
	"fmt"
	"math"
)

// TODO: implement the function computeQuadraticFormula

func computeDiscriminant(a, b, c float64) float64 {
	return math.Pow(b, 2) - 4*a*c
}

func computeQuadraticFormula(a, b, c float64) []float64 {
	D := computeDiscriminant(a, b, c)

	if D > 0 {
		x1 := (-b + math.Sqrt(D)) / (2 * a)
		x2 := (-b - math.Sqrt(D)) / (2 * a)
		return []float64{x1, x2}

	} else if D == 0 {
		x := -b / (2 * a)
		return []float64{x}

	} else {
		return []float64{}
	}
}

func main() {
	// TODO: call the function computeQuadraticFormula
	fmt.Println("Fall 1:", computeQuadraticFormula(3, 4, 1))
	// Erwartet: [-0.333... -1]

	fmt.Println("Fall 2:", computeQuadraticFormula(2, 4, 2))
	// Erwartet: [-1]

	fmt.Println("Fall 3:", computeQuadraticFormula(3, 4, 2))
	// Erwartet: []
}
