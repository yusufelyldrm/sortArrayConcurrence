package main

import (
	"fmt"
	"sort"
	"sync"
)

var sizeAnArray int
var UsersArray []int

func main() {
	getUsersArray()

	var wg sync.WaitGroup

	// Split the array into 4 subarrays
	partitionSize := sizeAnArray / 4
	partitions := make([][]int, 4)
	for i := 0; i < 4; i++ {
		wg.Add(1)
		start := i * partitionSize
		end := start + partitionSize
		if i == 3 {
			end = sizeAnArray
		}
		partition := UsersArray[start:end]
		partitions[i] = partition
		go func(p []int) {
			fmt.Println("Sorting subarray : ", p)
			// Sort the subarray
			sort.Ints(p)
			wg.Done()
		}(partition)
	}

	// Wait for all goroutines to finish
	wg.Wait()

	// Merge all 4 sorted subarrays into one large sorted array
	merged := merge(partitions)
	fmt.Println("The sorted array is: ", merged)

}

// "merge" function merges all 4 sorted subarrays into one large sorted array
func merge(arrays [][]int) []int {
	mergeTwo := func(left []int, right []int) []int {
		merged := make([]int, len(left)+len(right))
		i, j, k := 0, 0, 0
		for i < len(left) && j < len(right) {
			if left[i] < right[j] {
				merged[k] = left[i]
				i++
			} else {
				merged[k] = right[j]
				j++
			}
			k++
		}
		for i < len(left) {
			merged[k] = left[i]
			i++
			k++
		}
		for j < len(right) {
			merged[k] = right[j]
			j++
			k++
		}
		return merged
	}

	// Merge all 4 sorted subarrays into one large sorted array
	merged := mergeTwo(arrays[0], arrays[1])
	merged = mergeTwo(merged, arrays[2])
	merged = mergeTwo(merged, arrays[3])
	return merged
}

// getUsersArray returns a slice from users
func getUsersArray() {
	fmt.Print("Enter the size of the array: (It has to divide by 4)")
	fmt.Scan(&sizeAnArray)
	if sizeAnArray%4 != 0 {
		fmt.Println("The size of the array has to divide by 4")
		getUsersArray()
	} else {
		fmt.Println("The size of the array is: ", sizeAnArray)
	}

	for i := 0; i < sizeAnArray; i++ {
		var user int
		fmt.Printf("Enter the %v. element of array : ", i+1)
		fmt.Scan(&user)
		UsersArray = append(UsersArray, user)
	}
	fmt.Println("The array is: ", UsersArray)
}
