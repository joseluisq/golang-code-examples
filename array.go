package main

import "fmt"

func main() {
  var total float64 = 0
  x := [5]float64{10, 20, 30, 40, 50}

  for _, value := range x {
    total += value
  }

  fmt.Println(total / float64(len(x)))
}

