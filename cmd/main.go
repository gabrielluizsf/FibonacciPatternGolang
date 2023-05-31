package main

import (
	"fmt"
	"math"
)

func maxSearch(f func(float64) float64, a, b float64, numIterations int) float64 {
	// Definir os valores iniciais
	ratio := float64(fibonacci(numIterations-2)) / float64(fibonacci(numIterations))
	x1 := a + (1-ratio) * (b - a)
	x2 := a + ratio * (b - a)
	fx1 := f(x1)
	fx2 := f(x2)

	// Iterar até alcançar a condição de parada
	for i := 0; i < numIterations-2; i++ {
		if fx1 > fx2 {
			b = x2
			x2 = x1
			fx2 = fx1
			ratio = float64(fibonacci(numIterations-i-2)) / float64(fibonacci(numIterations-i-1))
			x1 = a + (1-ratio) * (b - a)
			fx1 = f(x1)
		} else {
			a = x1
			x1 = x2
			fx1 = fx2
			ratio = float64(fibonacci(numIterations-i-1)) / float64(fibonacci(numIterations-i))
			x2 = a + ratio * (b - a)
			fx2 = f(x2)
		}
	}

	// Retornar o máximo encontrado
	if fx1 > fx2 {
		return x1
	}
	return x2
}

func main() {
	// Definir a função a ser otimizada
	f := func(x float64) float64 {
		return -math.Pow(x-2, 2) + 5
	}

	// Definir o intervalo de busca e o número de iterações
	a := 0.0
	b := 4.0
	numIterations := 10

	// Realizar a busca pelo valor máximo
	maximum := maxSearch(f, a, b, numIterations)

	fmt.Printf("O valor máximo encontrado é %.2f\n", maximum)
}
