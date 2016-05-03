package main

import "fmt"

func main() {
  arr := []float64 {10,20,30,40,50}
  x := arraySum(arr)
  fmt.Println(x)
}

func arraySum(values []float64) float64 {
  total := 0.0

  for _, value := range values {
    total += value
  }

  return total
}

