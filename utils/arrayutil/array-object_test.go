package arrayutil

import (
	"fmt"
	"testing"
)

func TestArrayList_Contains(t *testing.T) {

	list := ArrayList{}
	list.Push(1)

	fmt.Println("array-object_test[11]>", list)
	fmt.Println("array-object_test[11]>", list.Contains(1))
}
