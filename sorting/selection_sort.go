package main

import "fmt"

func main() {
	unsorted_array := []int{43, 62, 18, 24, 87, 41, 20, 89, 76}

	swap := false
	for i := 0; i < len(unsorted_array); i++ {
		idx := i
		for j := i; j < len(unsorted_array); j++ {
			if unsorted_array[idx] > unsorted_array[j] {
				idx = j
				swap = true
			}
		}

		if !swap {
			fmt.Printf("Nothing to change.\n")
		}

		unsorted_array[i], unsorted_array[idx] = unsorted_array[idx], unsorted_array[i]
	}

	fmt.Printf("%v\n", unsorted_array)
}
