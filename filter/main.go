package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 3, 3, 3}
	filter(nums, func(i int) bool {
		if i < 10 {
			return true
		}
		return false
	})
	fmt.Println(nums)
}

func filter(nums []int, callback func(int) bool) []int {
	result := []int{}
	for _, v := range nums {
		if callback(v) {
			result = append(result, v)
		}
	}
	return result
}
