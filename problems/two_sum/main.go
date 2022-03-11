package main

import "fmt"

func main() {
	numbers := []int{3, 2, 4}
	fmt.Println(twoSum(numbers, 6))
}

// Faster Version (0ms)
func twoSum(nums []int, target int) []int {
	values := make(map[int]int, len(nums))
	for index, num := range nums {
		if indx, ok := values[target-num]; ok {
			return []int{index, indx}
		}
		values[num] = index
	}
	return []int{}
}

// Medium version (4ms)
// func twoSum(nums []int, target int) []int {
// 	values := make(map[int]int)
// 	for index, num := range nums {
// 		if indx, ok := values[target-num]; ok {
// 			return []int{index, indx}
// 		}
// 		values[num] = index
// 	}
// 	return []int{}
// }

// Slower version (40ms)
// func twoSum(nums []int, target int) []int {
// 	for i := 0; i < len(nums)-1; i++ {
// 		for j := i + 1; j < len(nums); j++ {
// 			if nums[i]+nums[j] == target {
// 				return []int{i, j}
// 			}
// 		}
// 	}
// 	return []int{0, 0}
// }
