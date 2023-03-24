package main

import (
	"fmt"
)

func primos(n int) bool {

	if n <= 1 {
		return false
	}

	for i := 2; i < n; i++ {
		if n%i == 0 {
			return false
		}

	}

	return true

}

func main() {

	var n int
	fmt.Println("elija un numero")
	_, err := fmt.Scan(&n)

	if err != nil {

		fmt.Println("Error al leer la opcion")
	}

	if primos(n) {
		fmt.Println("Es primo")
	} else {
		fmt.Println("no es primo")
	}

}
