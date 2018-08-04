package RemoveDuplicates

func removeDuplicates(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}
	l := 1
	for i, num := range nums {
		if i > 0 && nums[l-1] != num {
			nums[l] = num
			l++
		}
	}
	return l
}
