func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	v := x
	order := 1
	for v >= 10 {
		v = v / 10
		order++
	}

	if order == 1 {
		return true
	}

	for i := 0; i < order/2; i++ {
		number1 := x / int(math.Pow10(i)) % 10
		number2 := x / int(math.Pow10(order-1-i)) % 10
		if number1 != number2 {
			return false
		}
	}
	return true
}