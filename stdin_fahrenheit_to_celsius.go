package main

import "fmt"

/*
 * Program to covert Fahreheit to Censius value
 */
func main() {
  var fahrenheit float64
  fmt.Print("Fahrenheit (°F):")

  fmt.Scanf("%f", &fahrenheit)
  celsius := (fahrenheit - 32) * 5/9
  fmt.Println("Celsius (°C):", celsius)
}

