package problem509

func fib(n int) int {
	if n <= 1 {
		return n
	}

	a,b := 1, 0

	for i := 1; i < n; i++ {
		save := b
		b = a + b
		a = save
	}

	return b + a
}