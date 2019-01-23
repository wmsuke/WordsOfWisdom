package repositores

import (
	"fmt"
	"testing"
)

func TestFindOne(t *testing.T) {
	v := wordRepository{}
	result, error := v.FindOne(1)
	if error != nil {
		t.Fatalf("failed test %#v", error)
	}
	fmt.Println(result)
}