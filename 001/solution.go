package s001

func twoSum(nums []int, target int) []int {
	hmap := make(map[int]int)
	for i, v := range nums {
		if val, ok := hmap[target-v]; ok {
			return []int{i, val}
		} else {
			hmap[v] = i
		}
	}

	return []int{-1, -1}
}

func TwoSum(nums []int, target int) []int {
	return twoSum(nums, target)
}
