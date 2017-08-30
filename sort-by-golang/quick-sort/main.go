package main

import "fmt"

func main() {
	unsortArray := []int{24, 5, 95, 16, 27, 31, 64, 98, 80, 48}
	fmt.Println("\nFinal sort array:", quickSort(unsortArray, 0, 9))
}

func quickSort(unsortArray []int, start int, end int) []int {
	if start < end {
		mid := partition(unsortArray, start, end)
		quickSort(unsortArray, start, mid-1)
		quickSort(unsortArray, mid+1, end)
	}
	return unsortArray
}

func partition(unsortArray []int, start, end int) int {
	pivot := unsortArray[end]
	pindex := 0
	for i := 0; i < end; i++ {
		if unsortArray[i] <= pivot {
			unsortArray[i], unsortArray[pindex] = unsortArray[pindex], unsortArray[i]
			pindex++
		}
	}
	unsortArray[pindex], unsortArray[end] = unsortArray[end], unsortArray[pindex]
	return pindex
}
