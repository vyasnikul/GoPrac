package main

import (
	"fmt"
)

func linearSearch(arr []int, key int) bool {
	for i, val := range arr {
		if val == key {
			fmt.Println(i, ":", val)
			return true
		}
	}
	return false
}

func binarySearch(arr []int, key int) bool {
	min := 0
	max := len(arr) - 1
	median := 0
	for min <= max {
		median = (min + max) / 2
		if arr[median] < key {
			min = median + 1
		} else {
			max = median - 1
		}
	}
	if min == len(arr) || arr[min] != key {
		return false
	}
	fmt.Println(median, ":", key)
	return true
}

func main() {
	xs := []int{95, 78, 46, 58, 45, 86, 99, 251, 320}
	fmt.Println("Input Array: ", xs)
	fmt.Println("Linear Search Output:", linearSearch(xs, 99))
	fmt.Println("Binary Search Output:", binarySearch(xs, 86))
}
