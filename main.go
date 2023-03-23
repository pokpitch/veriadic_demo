package main

import "fmt"

func main() {
	s := []string{"x", "y", "z"}
	veriadicString("a", "b", "c")
	veriadicString(s...)
}

func veriadicString(v ...string) {
	for i := 0; i < len(v); i++ {
		fmt.Println(v[i])
	}
}
