package algorithms

func HashMap(str string) int {
	runes := []rune(str)
	keys := make([]rune, 0)
	hm := make(map[rune][2]int)

	for i, r := range runes {
		val := hm[r]
		if val[1] == 0 {
			hm[r] = [2]int{i, 1}
			keys = append(keys, r)
		} else {
			val[1] += 1
			hm[r] = val
		}
	}

	for _, k := range keys {
		if hm[k][1] == 1 {
			return hm[k][0]
		}
	}

	return -1
}
