package main

func minWindow(s string, t string) string {

	if len(t) > len(s) {
		return ""
	}

	targetFreqMap := make(map[byte]int)
	for i := range t {
		targetFreqMap[t[i]]++
	}

	from, to := 0, len(s)+1
	r := 0
	matchCount := 0

	for l := range s {
		for r < len(s) && matchCount < len(t) {
			if _, exists := targetFreqMap[s[r]]; exists {
				targetFreqMap[s[r]]--
				if targetFreqMap[s[r]] >= 0 {
					matchCount++
				}
			}
			r++
		}

		if matchCount == len(t) {
			if r-l < to-from {
				to = r
				from = l
			}
		}

		// Shrink the window
		if _, exists := targetFreqMap[s[l]]; exists {
			targetFreqMap[s[l]]++
			if targetFreqMap[s[l]] > 0 {
				matchCount--
			}
		}
	}

	if to == len(s)+1 {
		return ""
	}

	return s[from:to]
}
