package main

import "fmt"

// TODO: implement the function convertCelsiusToFahrenheit
func convertCelsiusToFahrenheit(c float64) float64 {
	return c*1.8 + 32
}

// TODO: implement the function convertFahrenheitToCelsius
func convertFahrenheitToCelsius(f float64) float64 {
	return (f - 32) / 1.8
}

func main() {
	// TODO: call the function convertCelsiusToFahrenheit
	// TODO: call the function convertFahrenheitToCelsius

	// --- Test 1: Gefrierpunkt ---
	c1 := 0.0
	f1 := convertCelsiusToFahrenheit(c1)
	fmt.Printf("%.2f°C sind %.2f°F\n", c1, f1) // Erwartet: 32.00

	// Rückrechnung
	back1 := convertFahrenheitToCelsius(f1)
	fmt.Printf("Zurückgerechnet: %.2f°C\n\n", back1) // Erwartet: 0.00

	// --- Test 2: Siedepunkt ---
	c2 := 100.0
	f2 := convertCelsiusToFahrenheit(c2)
	fmt.Printf("%.2f°C sind %.2f°F\n", c2, f2) // Erwartet: 212.00

	// Rückrechnung
	back2 := convertFahrenheitToCelsius(f2)
	fmt.Printf("Zurückgerechnet: %.2f°C\n\n", back2) // Erwartet: 100.00

	// --- Test 3: Ein krummer Wert (z.B. Raumtemperatur) ---
	c3 := 23.5
	f3 := convertCelsiusToFahrenheit(c3)
	fmt.Printf("%.2f°C sind %.2f°F\n", c3, f3) // Erwartet: 74.30

	// Rückrechnung
	back3 := convertFahrenheitToCelsius(f3)
	fmt.Printf("Zurückgerechnet: %.2f°C\n", back3) // Erwartet: 23.50
}
