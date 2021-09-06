package s033

func search(nums []int, target int) int {
	bot, top := 0, len(nums)-1

	for top-bot > 1 {
		mid := (top-bot)/2 + bot
		if nums[mid] == target {
			return mid
		}

		regCaseL := nums[bot] < nums[mid] && target >= nums[bot] && target <= nums[mid]
		pivCaseL := nums[bot] > nums[mid] && (target <= nums[mid] || target >= nums[bot])
		regCaseR := nums[mid] < nums[top] && target <= nums[top] && target >= nums[mid]
		pivCaseR := nums[mid] > nums[top] && (target >= nums[mid] || target <= nums[top])

		if regCaseL || pivCaseL {
			top = mid
		} else if regCaseR || pivCaseR {
			bot = mid
		} else {
			break
		}
	}

	if nums[bot] == target {
		return bot
	}
	if nums[top] == target {
		return top
	}

	return -1
}
