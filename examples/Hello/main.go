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

func interpolationSearch(arr []int, key int) int {
	Length := len(arr)
	min, max := arr[0], arr[Length-1]
	low, high := 0, Length-1
	for {
		if key < min {
			return low
		}
		if key > max {
			return Length
		}
		// Making guess of location
		guess := 0
		if high == low {
			guess = high
		} else {
			size := high - low
			offset := int(float64(size-1) * (float64(key-min) / float64(max-min)))
			guess = low + offset
		}
		// If element is found
		if arr[guess] == key {
			fmt.Println(guess, ":", key)
			return guess
		}
		// if guess is high, go lower direction or vice versa
		if arr[guess] > key {
			high = guess - 1
			max = arr[high]
		} else {
			low = guess + 1
			min = arr[low]
		}
	}
}

func quickSort(arr []int) []int {
	Length := len(arr)
	if Length < 2 {
		return arr
	}
	low, high := 0, Length-1
	pivot := int((Length - 1) / 2)
	arr[pivot], arr[high] = arr[high], arr[pivot]
	for index := range arr {
		if arr[index] < arr[high] {
			arr[low], arr[index] = arr[index], arr[low]
			low++
		}
	}
	arr[low], arr[high] = arr[high], arr[low]
	quickSort(arr[:low])
	quickSort(arr[low+1:])
	return arr
}

func main() {
	xs := []int{95, 78, 46, 58, 45, 86, 99, 251, 320}
	fmt.Println("<!DOCTYPE html>")
	fmt.Println("<html>")
	fmt.Println("<head>My First Web Page in Go</head>")
	fmt.Println("<body>")
	fmt.Println("Input Array: ", xs)
	fmt.Println("Linear Search Output:", linearSearch(xs, 99))
	fmt.Println("Binary Search Output:", binarySearch(xs, 86))
	fmt.Println("Interpolation Search Output:", interpolationSearch(xs, 251))
	fmt.Println("Quick Sort : ", quickSort(xs))
	fmt.Println("</body>")
	fmt.Println("</html>")
}
