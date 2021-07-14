package _78

func isBadVersion(version int) bool

func firstBadVersion(n int) int {
	left := 0
	right := n
	v := 0
	for ; left <= right; {
		idx := (left + right) / 2
		if isBadVersion(idx) {
			right = idx - 1
			v = idx
		} else {
			left = idx + 1
		}
	}

	return v
}
