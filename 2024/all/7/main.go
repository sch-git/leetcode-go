package main

func main() {

}

func reverse(x int) int {
	var res = 0
	for x != 0 {
		tmp := x % 10
		// check
		if res > 214748364 || (res == 214748364 && tmp > 7) {
			return 0
		}

		if res < -214748364 || (res == -214748364 && tmp > 8) {
			return 0
		}

		res = res*10 + tmp
		x = x / 10

	}
	return res
}
