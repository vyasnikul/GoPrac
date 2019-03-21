package main

import "fmt"

// PrintEvenNumbers : Printing Even numbers using for loop
func PrintEvenNumbers(x int) {
	// Using For
	for i := 0; i <= x; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}
}

// CheckPrime : Check the number is prime or not.
func CheckPrime(x int) {
	count := 0
	for i := 1; i < x/2; i++ {
		if x%i == 0 {
			count++
		}
	}
	if count == 1 {
		fmt.Println(x, " is prime.")
	} else {
		fmt.Println(x, " is not prime.")
	}
}

// findKIndex : find index of elements leading to value k
func findKIndex(arr []int, k int) (int, int) {
	index1, index2 := 0, 0
	// j := len(arr) - 1
	// for idx, val := range arr {
	// 	if val+arr[j] == k {
	// 		index1 = idx
	// 		index2 = j
	// 	} else {
	// 		fmt.Println("Hi")
	// 		j--
	// 	}
	// }
	for idx1, val1 := range arr {
		for idx2, val2 := range arr[idx1+1:] {
			if val1+val2 == k {
				index1 = idx1
				index2 = idx1 + idx2 + 1
			}
		}
	}
	return index1, index2
}

func main() {
	x := 0
	xs := []int{2, 4, 5, 6, 7, 8, 10, 10}
	fmt.Println("Enter the limit for Even numbers:")
	fmt.Scanf("%d", &x)
	fmt.Println("===========RESULT===============")
	PrintEvenNumbers(x)
	fmt.Println()
	CheckPrime(x)
	fmt.Println()
	fmt.Println(findKIndex(xs, x))
}
