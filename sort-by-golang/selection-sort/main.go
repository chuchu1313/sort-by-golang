package main

import (
	"fmt"
)

func main() {
	unsortArray := []int{3, 3, 19, 88, 97, 33, 25, 92, 86, 20}
	fmt.Println("\nFinal sort array:", selectionSort(unsortArray))
}

func selectionSort(unsortArray []int) []int {
	for i := 0; i < len(unsortArray); i++ {
		min := i
		for j := i + 1; j < len(unsortArray); j++ {
			if unsortArray[j] < unsortArray[min] {
				min = j
			}
		}
		unsortArray[i], unsortArray[min] = unsortArray[min], unsortArray[i]
		fmt.Println(unsortArray)
	}
	return unsortArray

}
