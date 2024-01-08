package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Codec struct {
}

// Encodes a list of strings to a single string.
func (codec *Codec) Encode(strs []string) string {
	builder := strings.Builder{}
	for i := range strs {
		builder.WriteString(fmt.Sprintf("%d", len(strs[i])))
		builder.WriteString("|")
		builder.WriteString(strs[i])
	}
	return builder.String()
}

// Decodes a single string to a list of strings.
func (codec *Codec) Decode(strs string) []string {
	res := make([]string, 0)

	i := 0
	for i < len(strs) {
		j := i
		for strs[j] != '|' {
			j++
		}

		len, _ := strconv.Atoi(strs[i:j])
		res = append(res, strs[j+1:j+1+len])
		i = j + 1 + len
	}

	return res
}

// Your Codec object will be instantiated and called as such:
// var codec Codec
// codec.Decode(codec.Encode(strs));
