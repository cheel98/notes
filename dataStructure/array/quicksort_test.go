package array

import (
	"fmt"
	"math/rand/v2"
	"testing"
)

func Test0(t *testing.T) {
	fmt.Printf("%v", quicksort([]int{-10, -8}))
}

func quicksort(arr []int) []int {
	n := len(arr)
	// 边界
	if n <= 1 {
		return arr
	}
	pivot := rand.IntN(n)
	arr[pivot], arr[n-1] = arr[n-1], arr[pivot]
	pivot = n - 1
	left, right := 0, n-2
	for left < right {
		for left < right && arr[left] < arr[pivot] {
			left++
		}
		for left < right && arr[right] > arr[pivot] {
			right--
		}
		arr[left], arr[right] = arr[right], arr[left]
	}
	arr[pivot], arr[left] = arr[left], arr[pivot]
	quicksort(arr[:left])
	quicksort(arr[left+1:])
	return arr
}
