package main

import (
	"fmt"
)

// BinarySearch procura por target em um slice ordenado.
// Retorna (índice, true) se encontrado, ou (-1, false) se não existir.
func BinarySearch(arr []int, target int) (int, bool) {
	low, high := 0, len(arr)-1

	for low <= high {
		mid := low + (high-low)/2

		if arr[mid] == target {
			return mid, true
		}
		if arr[mid] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1, false
}

func main() {
	arr := []int{0, 1, 1, 2, 3, 5, 8, 13, 21, 34}
	target := 21

	idx, found := BinarySearch(arr, target)
	if found {
		fmt.Printf("Elemento %d encontrado no índice %d\n", target, idx)
	} else {
		fmt.Printf("Elemento %d não encontrado\n", target)
	}
}
