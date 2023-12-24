package main

func longestValidParentheses(s string) int {
	if s == "" {
		return 0
	}

	open := 0
	close := 0
	res := 0

	for i := range s {
		if s[i] == '(' {
			open++
		} else {
			close++
		}

		if open == close {
			res = max(res, open*2)
			continue
		}

		if close > open {
			close, open = 0, 0
		}
	}

	open = 0
	close = 0

	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '(' {
			open++
		} else {
			close++
		}

		if open == close {
			res = max(res, open*2)
			continue
		}

		if open > close {
			close, open = 0, 0
		}
	}

	return res
}
