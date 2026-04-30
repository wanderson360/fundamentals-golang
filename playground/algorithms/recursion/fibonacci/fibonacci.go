package main

import (
	"errors"
	"fmt"
)

// Versão Recursiva: Retorna o n-ésimo termo da sequência - Focada na elegância teórica e estudo de stack
func FibonacciRecursion(n int) int {
	if n <= 1 {
		return n
	}
	return FibonacciRecursion(n-1) + FibonacciRecursion(n-2)
}

// FibonacciSeq gera a sequência completa até n elementos usando recursão.
func FibonacciSeq(n int) []int {
	seq := make([]int, n)
	for i := 0; i < n; i++ {
		seq[i] = FibonacciRecursion(i)
	}
	return seq
}

// Versão Iterativa: Focada em performance e economia de memória - (versão mundo real)
func Fibonacci(n int) ([]int, error) {
	if n < 0 {
		return nil, errors.New("n deve ser inteiro >= 0")
	}

	if n == 0 {
		return []int{}, nil
	}
	if n == 1 {
		return []int{0}, nil
	}

	seq := make([]int, n)
	seq[0], seq[1] = 0, 1

	for i := 2; i < n; i++ {
		seq[i] = seq[i-1] + seq[i-2]
	}

	return seq, nil
}

func main() {
	seq, err := Fibonacci(10)
	if err != nil {
		fmt.Println("Erro:", err)
		return
	}
	fmt.Println(seq)              // Saída: [0 1 1 2 3 5 8 13 21 34]
	fmt.Println(FibonacciSeq(10)) // Saída: [0 1 1 2 3 5 8 13 21 34]
}
