package main

func canPlaceFlowers(flowerbed []int, n int) bool {

	for i := 0; i < len(flowerbed); i++ {
		if flowerbed[i] == 1 {
			continue
		}

		leftEmpty, rightEmpty := false, false

		if i == 0 || flowerbed[i-1] == 0 {
			leftEmpty = true
		}

		if i == len(flowerbed)-1 || flowerbed[i+1] == 0 {
			rightEmpty = true
		}

		if leftEmpty && rightEmpty {
			flowerbed[i] = 1
			n--
		}

		if n <= 0 {
			return true
		}
	}

	return n <= 0
}
