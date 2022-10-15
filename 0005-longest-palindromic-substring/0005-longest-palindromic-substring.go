func longestPalindrome(s string) string {
	result, resLen := "", 0
	for idx, _ := range s {
		leftPtr, rightPtr := idx, idx
		result, resLen = expandAround(s, result, resLen, leftPtr, rightPtr)
		// even len
		leftPtr, rightPtr = idx, idx+1
		result, resLen = expandAround(s, result, resLen, leftPtr, rightPtr)
	}
	return result
}

func expandAround(s, result string, resLen, leftPtr, rightPtr int) (string, int) {
	for leftPtr >= 0 && rightPtr < len(s) && s[leftPtr] == s[rightPtr] {
		currLength := /* distance between the 2 ptrs */ (rightPtr - leftPtr + 1)
		if currLength > resLen {
			resLen = currLength
			result = s[leftPtr : rightPtr+1]
		}
		// fan out
		// keep going left
		leftPtr--
		// keep going right on right
		rightPtr++
	}
	return result, resLen
}