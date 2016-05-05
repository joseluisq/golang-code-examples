package main

import "fmt"

func main() {
	x := map[string]string{
		"go": "GoLang",
		"js": "Javascript",
		"py": "Python",
		"rb": "Ruby",
	}

	fmt.Println(x["go"])

	y := make(map[int]int)
	y[1] = 30
	fmt.Println(y[1])

	name, ok := x["js"]
	fmt.Println(name, ok)

	if val, ok := x["py"]; ok {
		fmt.Println(val, ok)
	}
}
