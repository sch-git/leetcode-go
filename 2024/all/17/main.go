package main

import (
	"fmt"
)

func main() {
	fmt.Println(letterCombinations("23"))
}

func letterCombinations(digits string) []string {
	if digits == "" {
		return nil
	}
	keyMap := map[int][]rune{
		2: {'a', 'b', 'c'},
		3: {'d', 'e', 'f'},
		4: {'g', 'h', 'i'},
		5: {'j', 'k', 'l'},
		6: {'m', 'n', 'o'},
		7: {'p', 'q', 'r', 's'},
		8: {'t', 'u', 'v'},
		9: {'w', 'x', 'y', 'z'},
	}

	stack := []string{""}
	for _, ch := range digits {
		knum := int(ch - '0')
		curLen := len(stack)
		for i := 0; i < curLen; i++ {
			curVal := stack[0]
			stack = stack[1:]
			for _, val := range keyMap[knum] {
				stack = append(stack, curVal+string(val))
			}
		}
	}
	return stack
}
