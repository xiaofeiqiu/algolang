package main

// Finds the longest palindromic substring in a given string.
func longestPalindrome(s string) string {
	var start, end int // Indices of the longest palindrome found

	for i := range s {
		left1, right1 := palindromeIndex(s, i, i)   // Odd length palindrome
		left2, right2 := palindromeIndex(s, i, i+1) // Even length palindrome

		if right1-left1 > end-start {
			start, end = left1, right1
		}
		if right2-left2 > end-start {
			start, end = left2, right2
		}
	}

	return s[start : end+1]
}

// Expands around the center indices and returns the bounds of the palindrome.
func palindromeIndex(s string, left, right int) (int, int) {
	for left >= 0 && right < len(s) && s[left] == s[right] {
		left--
		right++
	}
	return left + 1, right - 1
}
