#  Algoritmo de Otimização Fibonacci
```go
  
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

  

```
Este código em Go demonstra um algoritmo de otimização utilizando a sequência de Fibonacci para encontrar o valor máximo de uma função dentro de um intervalo especificado.

# Função Fibonacci

```go
func fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

```

A função fibonacci calcula a sequência de Fibonacci de forma recursiva. Ela recebe um número inteiro n como entrada. Essa função serve como auxiliar para o algoritmo de otimização.


#  Função maxSearch

A função maxSearch implementa o algoritmo de otimização utilizando a sequência de Fibonacci. Ela recebe os seguintes parâmetros:

    f (função): A função a ser otimizada. Ela deve aceitar um parâmetro do tipo float64 e retornar um valor do tipo float64.
    a (float64): O limite inferior do intervalo de busca.
    b (float64): O limite superior do intervalo de busca.
    numIterations (int): O número de iterações a serem realizadas na busca.

O algoritmo inicializa as variáveis necessárias e itera utilizando a sequência de Fibonacci para encontrar o valor máximo. Ele utiliza a função f fornecida para avaliar a função em cada ponto e atualiza o intervalo de busca de acordo.

O algoritmo de busca do máximo termina quando o número de iterações especificado é alcançado. Ele retorna o valor máximo encontrado.

# Função Principal
```go
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
```

A função main demonstra o uso do algoritmo de otimização. Ela define a função f a ser otimizada, o intervalo de busca [a, b] e o número de iterações. Em seguida, realiza a busca do máximo utilizando a função maxSearch e imprime o resultado.

Você pode modificar a função f, o intervalo de busca e o número de iterações de acordo com o seu problema específico. O algoritmo irá encontrar o valor máximo da função dentro do intervalo fornecido.

Sinta-se à vontade para explorar e experimentar diferentes funções e parâmetros de busca para otimizar diversos problemas utilizando o algoritmo de otimização Fibonacci.

# Pyramid Fibonacci
```go
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
```

Este código em Go exibe uma pirâmide de números de Fibonacci no terminal. Cada linha da pirâmide representa os primeiros números de Fibonacci até aquele ponto.

A função pyramid_fibonacci recebe um parâmetro numLinhas que representa o número de linhas da pirâmide a serem exibidas. Para cada linha, a função calcula o número de espaços em branco a serem exibidos antes dos números de Fibonacci e o número de elementos a serem exibidos naquela linha.

Em seguida, a função exibe os espaços em branco utilizando a função strings.Repeat para repetir o caractere de espaço em branco várias vezes. Após isso, ela itera sobre os números de Fibonacci correspondentes aquela linha e os exibe no terminal, separados por espaço.

Para utilizar essa função e exibir a pirâmide, você pode chamá-la no main do seu programa, passando o número desejado de linhas como argumento.

# Exemplo de uso:

```go

func main() {
	numLinhas := 5
	pyramid_fibonacci(numLinhas)
}

```

- Rode este comando no terminal:

```bash
go run *.go
```

Neste exemplo, a função pyramid_fibonacci será chamada para exibir uma pirâmide com 5 linhas.