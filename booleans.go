package main

import "fmt"

func main() {
  fmt.Println("true && true = ", true && true)
  fmt.Println("true && false = ", true && false)
  fmt.Println("true || false = ", true || true)
  fmt.Println("true || false = ", true || false)
  fmt.Println("false || false = ", false || false)
  fmt.Println("false && false = ", false && false)
  fmt.Println("!true = ", !true)
  fmt.Println("!false = ", !false)
}

