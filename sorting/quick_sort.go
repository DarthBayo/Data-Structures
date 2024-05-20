package main

import "fmt"

func partition(arr *[]int, low, high int) int {
	pivot := (*arr)[low]

	high_idx := high
	for i := high; i >= low; i-- {
		if (*arr)[i] > pivot {
			(*arr)[i], (*arr)[high_idx] = (*arr)[high_idx], (*arr)[i]
			high_idx--
		}
	}

	(*arr)[low], (*arr)[high_idx] = (*arr)[high_idx], (*arr)[low]
	return high_idx
}

func quickSort(arr *[]int, low, high int) {
	if low < high {
		pivot_index := partition(arr, low, high)

		quickSort(arr, low, pivot_index-1)
		quickSort(arr, pivot_index+1, high)
	}
}

func main() {
	unsorted_array := []int{43, 62, 18, 24, 87, 41, 20, 89, 76}

	quickSort(&unsorted_array, 0, len(unsorted_array)-1)

	fmt.Printf("%v\n", unsorted_array)
}
