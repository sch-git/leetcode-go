package _1816

import "strings"

// simple
func truncateSentence(s string, k int) string {
	slist := strings.Split(s," ")
	return strings.Join(slist[:k]," ")
}
