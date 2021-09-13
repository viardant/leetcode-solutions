package s1189

func maxNumberOfBalloons(text string) int {
    //store := [0,0,0,0,0]
    store := map[rune]int{'b': 0, 'a': 0, 'l': 0, 'o': 0, 'n': 0}
    for _, c := range text {
        switch c {
        case 'b', 'a', 'n':
            store[c] += 2
        case 'l', 'o':
            store[c] += 1
        }
    }
        
    min := store['b']
    for _, v := range store {
        if v < min {
            min = v
        }
    }
    
    return min / 2
}

func MaxNumberOfBalloons(text string) int {
	return maxNumberOfBalloons(text) 
}
