package main

import "strings"

func simplifyPath(path string) string {
	if path == "" {
		return ""
	}

	parts := strings.Split(path, "/")

	stack := make([]string, 0)
	for _, v := range parts {

		switch v {
		case "", ".":
			continue
		case "..":
			{
				if len(stack) > 0 {
					stack = stack[:len(stack)-1]
				}
			}
		default:
			stack = append(stack, v)
		}
	}

	return "/" + strings.Join(stack, "/")
}
