package repositores

import (
	"fmt"
	"testing"
)

func TestFindOne(t *testing.T) {
	var s = NewWordRepository()
	result, error := s.FindOne(1)
	// v := wordRepository{}
	// result, error := v.FindOne(1)
	if error != nil {
		t.Fatalf("failed test %#v", error)
	}
	fmt.Println(result)
}
