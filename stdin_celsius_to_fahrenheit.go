package main

import "fmt"

/*
 * Program to covert Censius to Fahrenheit value
 */
func main() {
	var celsius float64
	fmt.Print("Celsius (°C): ")
	fmt.Scanf("%f", &celsius)

	fahrenheit := (celsius * 9 / 5) + 32
	fmt.Println("Fahrenheit (°F):", fahrenheit)
}
