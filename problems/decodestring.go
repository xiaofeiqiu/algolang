package main

import "strings"

type E struct {
	Count int
	Str   string
}

func decodeString(s string) string {

	if s == "" {
		return ""
	}

	// track repeated count
	count := 0

	// track current string
	currStr := ""

	stack := make([]E, 0)

	for _, ch := range s {
		// Build count
		if ch >= '0' && ch <= '9' {
			count = count*10 + int(ch-'0')
			continue
		}

		// Build string
		if ch >= 'a' && ch <= 'z' {
			currStr += string(ch)
			continue
		}

		// new recurive starts, push string and count to stack
		// reset count and str
		if ch == '[' {
			stack = append(stack, E{Count: count, Str: currStr})
			count = 0
			currStr = ""
			continue
		}

		// current stack completed, calculate and pop the result
		// append repeated string to the string we built previousely
		// and assign it to currStr
		if ch == ']' {
			e := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			currStr = e.Str + strings.Repeat(currStr, e.Count)
		}
	}

	return currStr
}
