package main

func smallestEqual(nums []int) int {
	l := len(nums)
	for i := 0; i < l; i++ {
		if i%10 == nums[i] {
			return i
		}
	}
	return -1
}

func main() {

}