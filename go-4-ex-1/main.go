package main

import "fmt"

// TODO: implement the function computeGrade
func computeGrade(gotPoints float64, maxPoints float64) float64 {
	// Formel: 1 + 5 * (Erreicht / Max)
	result := 1.0 + 5.0*(gotPoints/maxPoints)
	return result
}

func main() {
	// TODO: call the function computeGrade
	grade1 := computeGrade(17.5, 28.0)
	fmt.Println(grade1) // 4.125

	grade2 := computeGrade(40.0, 40.0)
	fmt.Println(grade2) // 6

	grade3 := computeGrade(0.0, 20.0)
	fmt.Println(grade3) // 1
}
