package main

import "fmt"

func producto(a, b int) int {
	resultado := 0
	for i := 0; i < b; i++ {
		resultado += a
	}
	return resultado
}

func main() {
	a := 5
	b := 6
	resultado := producto(a, b)
	fmt.Printf("%d x %d = %d\n", a, b, resultado)
}
