package s1629

func slowestKey(releaseTimes []int, keysPressed string) byte {
    maxT, maxL := releaseTimes[0], keysPressed[0]
    for i, t := range releaseTimes {
        if i == 0 {
            continue
        }
        if t - releaseTimes[i-1] > maxT {
            maxT = t - releaseTimes[i-1]
            maxL = keysPressed[i]

        } else if t - releaseTimes[i-1] == maxT && keysPressed[i] > maxL {
                maxL = keysPressed[i]
        }
    }
    return maxL
}

func SlowestKey(releaseTimes []int, keysPressed string) byte {
	return slowestKey(releaseTimes, keysPressed)
}
