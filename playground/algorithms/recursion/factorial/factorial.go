package main

import "fmt"

// Versão Recursiva: Focada na elegância teórica e estudo de stack
func FactorialRecursive(n int) int {
	if n <= 1 {
		return 1
	}
	return n * FactorialRecursive(n-1)
}

// Versão Iterativa: Focada em performance e economia de memória - (versão mundo real)
func FactorialIterative(n int) int {
	if n < 0 {
		return -1
	}
	res := 1
	for i := 2; i <= n; i++ {
		res *= i
	}
	return res
}

func main() {
	val := 5
	fmt.Printf("Fatorial de %d (Recursivo): %d\n", val, FactorialRecursive(val))
	fmt.Printf("Fatorial de %d (Iterativo): %d\n", val, FactorialIterative(val))
}
