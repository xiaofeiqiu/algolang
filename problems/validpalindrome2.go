package main

func validPalindrome(s string) bool {
	if s == "" {
		return true
	}

	left, right := 0, len(s)-1

	for left < right {
		if s[left] != s[right] {
			return isPalindrome11(s, left+1, right) || isPalindrome11(s, left, right-1)
		}
		left++
		right--
	}

	return true
}

func isPalindrome11(s string, left, right int) bool {
	for left < right {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}

	return true
}
