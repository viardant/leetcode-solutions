package s003

func lengthOfLongestSubstring(s string) int {
    buffer := make(map[rune]int)
    max := 0
    trailing, leading := 0, -1

    for i, v := range s {
        if val, ok := buffer[v]; ok { // collision: rune already used 
            if leading - trailing +1 > max {
                max = leading - trailing + 1
            }
            // flush leftovers
            for j := trailing; j<=val; j++ {
                delete(buffer, rune(s[j]))
            }
            // cut substring
            trailing, leading = val+1, i
            // update buffer
            buffer[v] = i
        } else { // push forward
            buffer[v], leading = i, i
        }
    }
    if leading - trailing +1 > max {
                max = leading - trailing + 1
            }
    return max // fixes len of max subset
}

func LengthOfLongestSubstring(s string) int {
	return lengthOfLongestSubstring(s)
}
