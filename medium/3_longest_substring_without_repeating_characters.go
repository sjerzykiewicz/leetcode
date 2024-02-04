func lengthOfLongestSubstring(s string) int {
	curr_set := map[rune]struct{}{}
	var curr_slice []rune
	max_len := 0
	for _, char := range s {
		if _, ok := curr_set[char]; ok {
			if len(curr_slice) > max_len {
				max_len = len(curr_slice)
			}
			for _, c := range curr_slice {
				curr_slice = curr_slice[1:]
				delete(curr_set, c)
				if c == char {
					break
				}
			}
		}
		curr_slice = append(curr_slice, char)
		curr_set[char] = struct{}{}
	}
	if len(curr_slice) > max_len {
		return len(curr_slice)
	}
	return max_len
}
