package main

func numIdenticalPairs(nums []int) int {
	pair := 0
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums); j++ {
			if nums[i] == nums[j] && i < j {
				pair++
			}
		}
	}
	return pair
}
