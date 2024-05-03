package main

import "fmt"

func main() {
	unsorted_array, value := []int{43, 62, 18, 24, 87, 41, 20, 89, 76}, 20

	for i, v := range unsorted_array {
		if v == value {
			fmt.Printf("%v\nValue: %d\nPosition: %d\nLoops: %d", unsorted_array, value, i, (i + 1))
		}
	}
}
