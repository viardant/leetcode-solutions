package d29082021

func minPatches(nums []int, n int) int {
	miss, i, added := 1, 0, 0
	for miss <= n {
		if i < len(nums) && nums[i] <= miss {
			miss += nums[i]
			i++
		} else {
			miss += miss
			added++
		}
	}
	return added
}

func MinPatches(nums []int, n int) int {
	return minPatches(nums, n)
}
