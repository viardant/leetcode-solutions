package d30082021

func maxCount(m int, n int, ops [][]int) int {
	minX, minY := m, n
	if m == 1 && n == 1 {
		return 1
	}
	for _, o := range ops {
		changed := false
		if o[0] < minX {
			minX = o[0]
			changed = true
		}
		if o[1] < minY {
			changed = true
			minY = o[1]
		}
		if changed {
			if minX*minY == 1 {
				break
			}
		}
	}
	return minX * minY
}

func MaxCount(m, n int, ops [][]int) int {
	return maxCount(m, n, ops)
}
