package main

import "fmt"

// Реализовать бинарный поиск встроенными методами языка.

func main() {
	array := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	res := binarySearch(array, 3)

	fmt.Println(res)

}

func binarySearch(arr []int, n int) int {
	left := 0
	right := len(arr) -1

	for left <= right {
		mid := (left + right) / 2
		middle := arr[mid]

		if middle < n {
			left = mid + 1
		} else if middle > n {
			right = mid - 1
		} else {
			return mid
		}
	}

	return -1
}