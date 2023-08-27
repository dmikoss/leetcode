func lengthOfLongestSubstring(s string) int {

	hashmap := map[byte]int{}

	if len(s) < 2 {
		return len(s)
	}

	length := 0
	l := 0
	for r := 0; r < len(s); r++ {
		c := s[r]
		pos, exists := hashmap[c]
		if exists && pos >= l {
			l = pos + 1
		} else {
			length = max(length, r-l+1)
		}
		hashmap[c] = r
	}
	return length
}

func max(v1, v2 int) int {
	if v1 > v2 {
		return v1
	}
	return v2
}