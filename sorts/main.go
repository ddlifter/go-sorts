package main

import (
	"fmt"
	"math/rand"
)

func main() {
	nums := make([]int, 5, 5)
	for i := 0; i < 5; i++ {
		nums[i] = rand.Intn(100)
	}
	fmt.Println(nums)

	fmt.Println("-------")

	heap(nums)

	fmt.Println("-------")

	fmt.Println(nums)
}

// Худший случай: O(n^2)
// Лучший случай: O(n)
// Средний случай: O(n^2)
func bubble(a []int) {
	for i := 1; i < len(a); i++ {
		var ok bool = true
		for j := 0; j < len(a)-i; j++ {
			if a[j] > a[j+1] {
				a[j], a[j+1] = a[j+1], a[j]
				ok = false
			}
		}
		if ok {
			break
		}
		fmt.Println(a)
	}
}

// Худший случай: O(n^2)
// Лучший случай: O(n)
// Средний случай: O(n^2)
func insert(a []int) {
	for i := 1; i < len(a); i++ {
		j := i

		for j > 0 {
			if a[j-1] > a[j] {
				a[j-1], a[j] = a[j], a[j-1]
			}
			j = j - 1
		}
		fmt.Println(a)
	}
}

// Худший случай: O(n^2)
func binaryInsert(a []int) {
	for i := 1; i < len(a); i++ {
		key := a[i]
		left := 0
		high := i - 1
		for left <= high {
			mid := (high + left) / 2
			if key < a[mid] {
				high = mid - 1
			} else {
				left = mid + 1
			}
		}
		for j := i - 1; j >= left; j-- {
			a[j+1] = a[j]
		}
		a[left] = key
		fmt.Println(a)
	}
}

// Худший случай: O(n^2)
// Лучший случай: O(n^2)
// Средний случай: O(n^2)
func selection(a []int) {
	for i := 0; i < len(a); i++ {
		min := i
		for j := i + 1; j < len(a); j++ {
			if a[j] < a[min] {
				min = j
			}
		}
		a[min], a[i] = a[i], a[min]
		fmt.Println(a)
	}
}

// Худший случай: O(nlogn)
// Лучший случай: O(nlogn)
// Средний случай: O(nlogn)
func heap(a []int) {

}
