package main

import "fmt"

// Реализовать быструю сортировку массива (quicksort) встроенными методами языка.

func main() {
	array := []int{3, 14, 1, 7, 9, 8, 11, 6, 4, 2}

	result := quickSort(array)

	fmt.Println(result)
}

func quickSort(arr []int) []int {
	
	if len(arr) < 2 {
		return arr
	}

	left := 0
	right := len(arr) - 1
	mid := (left + right) / 2

	arr[mid], arr[right] = arr[right], arr[mid]

	for i := range arr {
		if arr[i] < arr[right] {
			arr[i], arr[left] = arr[left], arr[i]
			left++
		}
	}

	arr[left], arr[right] = arr[right], arr[left]

	quickSort(arr[:left])
	quickSort(arr[left + 1:])

	return arr

}
