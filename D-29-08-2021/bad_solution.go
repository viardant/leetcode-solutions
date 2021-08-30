package d29082021

/*
#V1
	I generate first all sums from subset. Then I keep adding the smallest non-generated untill I can manage to generate the whole span
	TOO SLOW
*/

func main() {
	minPatches([]int{1, 2, 2, 5}, 10)
	minPatches([]int{1, 2, 2}, 5)
	minPatches([]int{1, 5, 10}, 20)
}

func minPatches_bad(nums []int, n int) int {
	generated := make(map[int]bool)
	var nextMin *int = new(int)
	*nextMin = 1
	iterate(nums, 0, &generated, n, nextMin)
	patches := []int{}
	for *nextMin <= n {
		toAdd := *nextMin
		for {
			(*nextMin)++
			if _, found := generated[*nextMin]; !found {
				break
			}
		}
		patches = append(patches, toAdd)
		for genNum := range generated {
			if genNum+toAdd <= n {
				generated[genNum+toAdd] = true
			}
			if genNum == *nextMin || genNum+toAdd == *nextMin {
				for {
					(*nextMin)++
					if _, found := generated[*nextMin]; !found {
						break
					}
				}
			}
		}
	}

	return len(patches)
}

func iterate(nums []int, sum int, g *map[int]bool, top int, nextMin *int) {
	taken := sum + nums[0]
	if taken <= top {
		(*g)[taken] = true // taken
		if taken == *nextMin {
			for {
				(*nextMin)++
				if _, found := (*g)[*nextMin]; !found {
					break
				}
			}
		}
	}
	if len(nums) >= 2 {
		iterate(nums[1:], taken, g, top, nextMin) // taken
		iterate(nums[1:], sum, g, top, nextMin)   // not taken
	}
}
