package s033

func slow_search(nums []int, target int) int {
	return binSearchRotated(0, len(nums)-1, target, nums)
}

func binSearchRotated(bot, top, target int, nums []int) int {
	// final case
	if top-bot <= 1 {
		if nums[bot] == target {
			return bot
		}
		if nums[top] == target {
			return top
		}
		return -1
	}

	mid := (top-bot)/2 + bot
	if nums[mid] == target {
		return mid
	}

	regCaseL := nums[bot] < nums[mid] && target >= nums[bot] && target <= nums[mid]
	pivCaseL := nums[bot] > nums[mid] && (target <= nums[mid] || target >= nums[bot])
	regCaseR := nums[mid] < nums[top] && target <= nums[top] && target >= nums[mid]
	pivCaseR := nums[mid] > nums[top] && (target >= nums[mid] || target <= nums[top])

	if regCaseL || pivCaseL {
		// go left
		return binSearchRotated(bot, mid, target, nums)
	}
	if regCaseR || pivCaseR {
		// go rigth
		return binSearchRotated(mid, top, target, nums)
	}

	return -1
}
