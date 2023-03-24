package main

import (
	"fmt"
)

func main() {

	var opcion int

	fmt.Println("menu de opciones")
	fmt.Println("opcion 1 bife y pure")
	fmt.Println("opcion 2 milanesa y papas")
	fmt.Println("opcion 3 sorrentinos")
	fmt.Println("opcion 4 hamburguesa")
	fmt.Println("porfavor eliga una opcion del 1 al 4")

	_, err := fmt.Scan(&opcion)

	if err != nil {

		fmt.Println("Error al leer la opcion")
	}

	switch opcion {

	case 1:
		fmt.Println("Usted eligio bife y pure")

	case 2:
		fmt.Println("Usted eligio milanesa y papas")

	case 3:
		fmt.Println("elegiste sorrentinos")

	case 4:
		fmt.Println("elegiste hamburguesa")

	default:
		fmt.Println("opcion incorrecta")

	}

}
