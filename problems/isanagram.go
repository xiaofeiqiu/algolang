package main

func isAnagram(s string, t string) bool {

	chMap := make(map[byte]int)
	for i := range s {
		chMap[s[i]]++
	}

	for i := range t {
		chMap[t[i]]--
	}

	for _, v := range chMap {
		if v != 0 {
			return false
		}
	}

	return true
}
