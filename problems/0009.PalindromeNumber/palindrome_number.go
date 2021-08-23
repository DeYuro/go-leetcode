package problem9

func isPalindrome(x int) bool  {
	if x < 0 || (x > 0 && x % 10 == 0) {
		return false
	}

	if x == 0  {
		return true
	}

	var r  int

	for r < x {
		r = r * 10 + x % 10
		x = x / 10
	}

	if r > x {
		return r / 10 == x
	}

	return x == r
}

func isPalindrome2(x int) bool  {
	t := x
	if x < 0 {
		return false
	}

	var r, d int

	for t != 0 {
		d = t%10
		r = r * 10 + d
		t /= 10
	}

	return r == x
}