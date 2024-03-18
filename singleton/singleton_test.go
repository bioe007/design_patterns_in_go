package singleton

import (
	"testing"
)

// TODO - This doesn't work because in go getting the singleton creates a new
// memory address where the pointer to the singleton is stored.
// func TestSingletonSameAddress(t *testing.T) {
// 	unique1 := GetSingleton()
// 	unique2 := GetSingleton()
//
// 	if &unique1 != &unique2 {
// 		t.Error("Not the same, Jim.")
// 	}
// }

func TestAddFour(t *testing.T) {
	GetSingleton().Add(2)
	GetSingleton().Add(2)

	sngl := GetSingleton()
	if sngl.Get() != 4 {
		t.Error("wrong values", sngl.Get())
	}
}
