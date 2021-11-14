package _677

import "testing"

func TestName(t *testing.T) {
	obj := Constructor()
	obj.Insert("apple", 3)
	obj.Insert("app", 2)
	obj.Sum("ap")
}
