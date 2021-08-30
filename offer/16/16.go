package _16

// middle
// 快速幂

func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	if n == 1 {
		return x
	}

	if n < 0 {
		x = 1 / x
		n = -n
	}

	res := 1.0
	for n > 0 {
		if n&1 == 1 {
			res *= x
			n--
		} else {
			x *= x
			n >>= 1
		}
	}

	return res
}
