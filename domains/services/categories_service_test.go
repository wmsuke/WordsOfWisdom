package services

import (
	"fmt"
	"testing"
)

func TestGet(t *testing.T) {
	ctrl := NewCategory()
	// result, error := ctrl.Get(1)
	// if error != nil {
	// 	t.Fatalf("failed test %#v", error)
	// }
	result := ctrl.Get(1)
	fmt.Println(result)
	// if result.Category != "偉人" {
	// 	t.Fatal("failed test")
	// }
}
