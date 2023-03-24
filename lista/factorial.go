package main

import (
	"fmt"
)

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	result := 1
	for i := 1; i <= n; i++ {
		result *= i
	}
	return result
}

func main() {
	// Calcular el factorial de 5 y mostrar el resultado

	var n int
	fmt.Println("elija un numero")
	_, err := fmt.Scan(&n)

	if err != nil {

		fmt.Println("Error al leer la opcion")
	}
	resultado := factorial(n)
	fmt.Printf("El factorial de %d es %d\n", n, resultado)
}
