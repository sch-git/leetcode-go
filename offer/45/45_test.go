package _45

import (
	"testing"
	"unsafe"
)

func TestName(t *testing.T) {
	//minNumber([]int{3,30,34,5,9})
	type stu struct {
		int8
		int16
		int32
	}
	s := stu{}
	t.Log(unsafe.Sizeof(s))
}
