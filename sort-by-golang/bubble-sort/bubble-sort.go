package main

import (
	"fmt"
)

func main() {
	unsortArray := []int{44, 11, 23, 95, 36, 92, 33, 58, 84, 34}
	fmt.Println("\nInitial:", unsortArray)
	fmt.Println("Final sort array:", bubbleSort(unsortArray))
}


func bubbleSort(unsortArray []int) []int {
	sortFlag := true
	for i := 0; i < len(unsortArray)-1 && sortFlag; i++ {
		sortFlag = false
		for j := 0; j < len(unsortArray)-i-1; j++ {
			if unsortArray[j] > unsortArray[j+1] {
				unsortArray[j], unsortArray[j+1] = unsortArray[j+1], unsortArray[j]
				sortFlag = true
			}

		}
		fmt.Println(unsortArray)
	}
	return unsortArray
}
