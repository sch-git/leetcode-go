package _02

import "strconv"

// simple
// 二进制加法

func addBinary(a string, b string) string {
	flag := 0
	res := ""
	i, j := len(a)-1, len(b)-1
	for i >= 0 && j >= 0 {
		numa, _ := strconv.Atoi(string(a[i]))
		numb, _ := strconv.Atoi(string(b[j]))

		tempRes := (numa + numb + flag) % 2
		flag = (numa + numb + flag) / 2
		res = strconv.Itoa(tempRes) + res

		i--
		j--
	}

	for ; i >= 0; i-- {
		numa, _ := strconv.Atoi(string(a[i]))
		tempRes := (numa + flag) % 2
		flag = (numa + flag) / 2
		res = string(rune(tempRes)) + res
	}

	for ; j >= 0; j-- {
		numb, _ := strconv.Atoi(string(b[j]))
		tempRes := (numb + flag) % 2
		flag = (numb + flag) / 2
		res = string(rune(tempRes)) + res
	}

	// 10,11 最终还有进位的情况
	if flag > 0 {
		res = strconv.Itoa(flag) + res
	}

	return res
}
