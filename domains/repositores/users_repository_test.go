package repositores

import (
	"fmt"
	"testing"
)

func TestIsOne(t *testing.T) {
	var v = NewUserRepository()
	result, error := v.IsOne("dVachSk3sK1NpN3dri0wp0KOh1gqsSZMODrqo7Zv3WF3kzgXWCAfMCwIN6gYCdHI")
	if error != nil {
		t.Fatalf("failed test %#v", error)
	}
	fmt.Println(result)
}
