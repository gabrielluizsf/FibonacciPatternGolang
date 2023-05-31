package main


import (
	"fmt"
	"strings"
)


func pyramid_fibonacci(numLinhas int){

	for i := 0; i < numLinhas; i++ {
		numEspacos := numLinhas - i - 1
		numElementos := i + 1

		// Exibir espaços em branco
		fmt.Print(strings.Repeat(" ", numEspacos))

		// Exibir números de Fibonacci
		for j := 0; j < numElementos; j++ {
			numFibonacci := fibonacci(j)
			fmt.Print(numFibonacci, " ")
		}

		fmt.Println()
	}
}