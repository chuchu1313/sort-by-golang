package main

import "fmt"

func main() {
	unsortArray := []int{24, 5, 95, 16, 27, 31, 64, 98, 80, 48}
	fmt.Println("\nFinal sort array:", mergeSort(unsortArray))
}

func mergeSort(unsortArray []int) []int {
	if len(unsortArray) < 2 {
		return unsortArray
	} else {
		mid := len(unsortArray) / 2
		nL := make([]int, mid)
		nR := make([]int, len(unsortArray)-mid)
		for i := 0; i < len(nL); i++ {
			nL[i] = unsortArray[i]
		}
		for i := 0; i < len(nR); i++ {
			nR[i] = unsortArray[mid+i]
		}
		mergeSort(nL)
		mergeSort(nR)
		merge(nL, nR, unsortArray)
		return unsortArray
	}

}

func merge(nL []int, nR []int, unsortArray []int) {
	indexO, indexL, indexR := 0, 0, 0
	for indexL < len(nL) && indexR < len(nR) {
		if nL[indexL] <= nR[indexR] {
			unsortArray[indexO] = nL[indexL]
			indexL++
		} else {
			unsortArray[indexO] = nR[indexR]
			indexR++
		}
		indexO++
	}
	if indexL == len(nL) {
		for indexR < len(nR) {
			unsortArray[indexO] = nR[indexR]
			indexR++
			indexO++
		}
	} else if indexR == len(nR) {
		for indexL < len(nL) {
			unsortArray[indexO] = nL[indexL]
			indexL++
			indexO++
		}
	}
}
