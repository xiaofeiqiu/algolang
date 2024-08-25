package main

func myPow(x float64, n int) float64 {
	if x == 0 {
		return 0
	}

	if n == 0 {
		return 1
	}

	if n < 0 {
		return 1.0 / myPow(x, -n)
	}

	half := myPow(x, n/2)
	if n%2 == 0 {
		return half * half
	}

	return x * half * half
}
