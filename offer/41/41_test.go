package _41

import (
	"sort"
	"testing"
)

func TestName(t *testing.T) {
	m := &maxHeap{IntSlice: make(sort.IntSlice, 0)}
	m.IntSlice = append(m.IntSlice,1)
	m.IntSlice = append(m.IntSlice,3)
	m.IntSlice = append(m.IntSlice,2)
	sort.Ints(m.IntSlice)
	t.Log(m.IntSlice)
}
