package main

import (
	"strings"
)

func main() {

}

func intToRoman(num int) string {
	var m = map[int]string{
		4:   "IV",
		9:   "IX",
		40:  "XL",
		90:  "XC",
		400: "CD",
		900: "CM",
	}
	if _, ok := m[num]; ok {
		return m[num]
	}

	chm := map[int]string{
		1:    "I",
		4:    "IV",
		5:    "V",
		9:    "IX",
		10:   "X",
		40:   "XL",
		50:   "L",
		90:   "XC",
		100:  "C",
		400:  "CD",
		500:  "D",
		900:  "CM",
		1000: "M",
	}
	nums := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	res := strings.Builder{}
	for num > 0 {
		for _, n := range nums {
			if num >= n {
				for times := num / n; times > 0; times-- {
					res.WriteString(chm[n])
				}
				num = num % n
			}
		}
	}
	return res.String()
}
