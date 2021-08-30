package s027

func removeElement(nums []int, val int) int {
	lo, hi := 0, len(nums)-1
	for lo <= hi {

		if nums[lo] == val {
			nums[lo], nums[hi] = nums[hi], nums[lo]
			hi--
		} else {
			lo++
		}
	}

	return lo
}
