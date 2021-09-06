package s565

func arrayNesting(nums []int) int {
    max := 1
    for i := range nums {
        if nums[i] >= 0 {
            res := 0
            for start := i; nums[start] >= 0; {
                res++
                hold := start
                start = nums[start]
                nums[hold] = -1
            }
            if res > max {
                max = res
            }
            nums[i] = -1
        }

    }
    return max
}

func ArrayNesting(nums []int) int {
	return arrayNesting(nums)
}
