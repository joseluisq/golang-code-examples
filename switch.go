package main

import "fmt"

func main() {
	for i := 0; i <= 3; i++ {
		switch i {
		case 1:
			fmt.Println("One")
		case 2:
			fmt.Println("Two")
		case 3:
			fmt.Println("Three")
		default:
			fmt.Println("Unknown Number")
		}
	}
}
