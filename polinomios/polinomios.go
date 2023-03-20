package main

import (
	"fmt"
)

func mostrarPolinomio(coeficientes []float64) {
	polinomio := ""
	for i := len(coeficientes) - 1; i >= 0; i-- {
		grado := len(coeficientes) - 1 - i
		termino := ""
		if coeficientes[i] != 0 {
			if grado == 0 {
				termino = fmt.Sprintf("%.1f", coeficientes[i])
			} else if grado == 1 {
				termino = fmt.Sprintf("%.1f X", coeficientes[i])
			} else {
				termino = fmt.Sprintf("%.1f X**%d", coeficientes[i], grado)
			}
			if polinomio == "" {
				polinomio = termino
			} else {
				polinomio = termino + " + " + polinomio
			}
		}
	}
	fmt.Println(polinomio)
}

func main() {
	coeficientes := []float64{3.0, 2.0, 1.0}
	mostrarPolinomio(coeficientes)
}
