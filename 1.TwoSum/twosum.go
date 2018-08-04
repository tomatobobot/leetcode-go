package TwoSum

func twoSum(nums []int, target int) []int {
	ts := make(map[int]int)
	for i, val := range nums {
		num := target - val
		if idx, ok := ts[num]; ok {
			return []int{idx, i}
		}
		ts[val] = i
	}
	return []int{}
}
