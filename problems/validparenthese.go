package main

func isValid(s string) bool {

	if s == "" {
		return true
	}

	// Map of matching bracket pairs.
	pairs := map[rune]rune{
		'(': ')',
		'{': '}',
		'[': ']',
	}

	// Stack for keeping track of brackets.
	stack := make([]rune, 0)

	for _, char := range s {
		// if is opening, push expected closing to stack
		if closing, isOpening := pairs[char]; isOpening {
			stack = append(stack, closing)
			continue
		}

		// if not opening, compare expected closing with char
		// if match found, pop expected closing
		if len(stack) != 0 && stack[len(stack)-1] == char {
			stack = stack[:len(stack)-1]
			continue
		}

		// if match not found, return false
		return false
	}

	// if all poped, meaning all closings and openings are matched
	return len(stack) == 0
}
