package main

import "fmt"

func main() {
	sorted_array, value := []int{8, 13, 20, 21, 30, 33, 49, 50, 52, 64}, 50

	low, mid, high := 0, 0, len(sorted_array)
	for i, _ := range sorted_array {
		mid = low + ((high - low) / 2)

		if sorted_array[mid] == value {
			fmt.Printf("%v\nValue: %d\nPosition: %d\nLoops: %d", sorted_array, value, mid, i)
			return
		}

		if sorted_array[mid] > value {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	return
}
