package collections

import (
	"fmt"
	"testing"
)

func Test_EmptyList(t *testing.T) {

	list := NewList[int]()

	if list == nil {
		t.Error("invaid collection")
	}
	fmt.Println(list)
}

func Test_ListWithData(t *testing.T) {

	list := NewListT(1, 2, 3, 4, 5, 6)

	if list == nil {
		t.Error("invaid collection")
	}
	fmt.Println(list)
}
