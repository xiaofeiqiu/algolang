package main

func minRemoveToMakeValid(s string) string {

	if s == "" {
		return s
	}

	stack := make([]int, 0)
	for i := range s {
		if s[i] == '(' {
			stack = append(stack, i)
			continue
		}

		if s[i] == ')' {
			if len(stack) > 0 && s[stack[len(stack)-1]] == '(' {
				stack = stack[:len(stack)-1]
				continue
			}
			stack = append(stack, i)
		}
	}

	toRemove := make(map[int]bool)
	for i := range stack {
		toRemove[stack[i]] = true
	}

	res := make([]byte, 0)
	for i := range s {
		if !toRemove[i] {
			res = append(res, s[i])
		}
	}
	return string(res)

}
