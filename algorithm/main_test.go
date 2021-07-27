package algorithm

import (
	_33 "leetcode-go/algorithm/day7/733"
	"testing"
)

func TestName(t *testing.T) {

	_33.FloodFill([][]int{
		{
			0, 0, 0,
		},
		{
			0, 1, 1,
		},
	}, 1, 1, 1)
}
