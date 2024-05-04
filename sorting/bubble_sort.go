package main

import "fmt"

func main() {
	unsorted_array := []int{43, 62, 18, 24, 87, 41, 20, 89, 76}

	swap := false
	for i := 0; i < len(unsorted_array); i++ {
		for j := 0; j < len(unsorted_array)-1; j++ {
			if unsorted_array[j] > unsorted_array[j+1] {
				unsorted_array[j], unsorted_array[j+1] = unsorted_array[j+1], unsorted_array[j]
				swap = true
			}
		}

		if !swap {
			fmt.Printf("Nothing to change.\n")
		}
	}

	fmt.Printf("%v\n", unsorted_array)
}
