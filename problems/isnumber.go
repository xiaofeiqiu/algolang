package main

import "strings"

func isNumber(s string) bool {
	if s == "" {
		return false
	}

	s = strings.ToLower(s)
	parts := strings.Split(s, "e")

	if len(parts) > 2 {
		return false
	}

	leftValid := false
	rightValid := true

	if len(parts) > 0 {
		leftValid = validate(parts[0], true)
	}

	if len(parts) == 2 {
		rightValid = validate(parts[1], false)
	}

	return leftValid && rightValid
}

func validate(s string, allowDot bool) bool {
	if s == "" {
		return false
	}

	if s[0] == '+' || s[0] == '-' {
		s = s[1:]
	}

	digitSeen := false
	dotSeen := false

	for i := range s {
		if s[i] >= '0' && s[i] <= '9' {
			digitSeen = true
			continue
		}

		if s[i] == '.' && allowDot && !dotSeen {
			dotSeen = true
			continue
		}

		return false
	}

	return digitSeen
}
