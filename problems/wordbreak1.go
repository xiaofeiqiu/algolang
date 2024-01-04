package main

func wordBreak(s string, wordDict []string) bool {
	if wordDict == nil {
		return false
	}

	n := len(s)
	dp := make([]bool, n+1)
	dp[0] = true

	for i := 1; i <= n; i++ {
		for _, word := range wordDict {
			l := len(word)
			if i < l {
				continue
			}

			if dp[i-l] && s[i-l:i] == word {
				dp[i] = true
				break
			}
		}
	}

	return dp[n]
}
