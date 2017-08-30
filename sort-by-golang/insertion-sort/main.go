package main

import (
	"fmt"
)

func main() {
	unsortArray := []int{93, 20, 73, 73, 36, 79, 64, 62, 37, 31}
	fmt.Println("\nInitial:", unsortArray)
	fmt.Println("\nFinal sort array:", insertionSort(unsortArray))
}

func insertionSort(unsortArray []int) []int {
	for i := 1; i < len(unsortArray); i++ {
		for j := i; j > 0; j-- {
			if unsortArray[j] < unsortArray[j-1] {
				unsortArray[j], unsortArray[j-1] = unsortArray[j-1], unsortArray[j]
			} else {
				break
			}
		}
		fmt.Println(unsortArray)
	}
	return unsortArray
}
