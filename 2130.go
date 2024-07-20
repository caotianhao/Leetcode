package main

type ListNode2130 struct {
	Val  int
	Next *ListNode2130
}

func pairSum(head *ListNode2130) (max int) {
	arr := make([]int, 0)
	for head != nil {
		arr = append(arr, head.Val)
		head = head.Next
	}
	left, right := 0, len(arr)-1
	for left < right {
		max = max2130(max, arr[left]+arr[right])
		left++
		right--
	}
	return
}

func max2130(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {

}
